package service

import (
	"context"
	"errors"
	"fmt"
	"kreditplus/src/app"
	"kreditplus/src/entity"
	"kreditplus/src/middleware"
	"kreditplus/src/tracer"
	"kreditplus/src/v1/contract"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

func (svc TransactionService) Create(ctx context.Context, req *contract.TransactionRequest) (contract.CreateTransactionResponse, error) {
	ctx, span := app.Tracer().Start(ctx, "CreateTransactionService")
	defer span.End()

	resp := contract.CreateTransactionResponse{}

	db := app.GormDB().WithContext(ctx)

	// Check API Key
	partnerCount := int64(0)
	err := db.Model(&entity.Partner{}).Where("api_key", req.APIKey).Count(&partnerCount).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	if partnerCount == 0 {
		err := errors.New("partner not found")
		tracer.RecordError(span, err)
		return resp, err
	}

	// Get Partner
	partner := entity.Partner{}
	err = db.Model(&entity.Partner{}).Where("api_key", req.APIKey).Find(&partner).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	if !partner.IsActive {
		err := errors.New("partner not active")
		tracer.RecordError(span, err)
		return resp, err
	}

	// Check if Order ID already exist
	lastTransactionCount := int64(0)
	err = db.Model(&entity.Transaction{}).Where("order_id", req.OrderID).Count(&lastTransactionCount).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if lastTransactionCount > 0 {
			lastTransaction := entity.Transaction{}
			err = tx.Order("transaction_date desc").Where("order_id", req.OrderID).First(&lastTransaction).Error
			if err != nil {
				tracer.RecordError(span, err)
				return err
			}

			if lastTransaction.Status == contract.TransactionComplete || lastTransaction.Status == contract.TransactionSettlement {
				err = errors.New("transaction already success")
				tracer.RecordError(span, err)
				return err
			}

			lastTransaction.Status = contract.TransactionCancelled
			err := tx.Model(&entity.Transaction{}).Where("id", lastTransaction.ID).Save(&lastTransaction).Error
			if err != nil {
				tracer.RecordError(span, err)
				return err
			}
		}

		txnID, err := uuid.NewV7()
		if err != nil {
			tracer.RecordError(span, err)
			return err
		}

		transaction := entity.Transaction{
			PartnerID:       partner.ID,
			Phone:           req.Phone,
			AssetName:       req.AssetName,
			AdminFee:        req.AdminFee,
			OTR:             req.OTR,
			Status:          contract.TransactionInit,
			TxnID:           txnID.String(),
			OrderID:         req.OrderID,
			TransactionDate: time.Now(),
		}

		err = tx.Model(&entity.Transaction{}).Create(&transaction).Error
		if err != nil {
			tracer.RecordError(span, err)
			return err
		}

		resp.TxnID = txnID.String()
		resp.OrderID = req.OrderID

		return nil
	})
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (svc TransactionService) Get(ctx context.Context, txnID string) (contract.GetTransactionResponse, error) {
	ctx, span := app.Tracer().Start(ctx, "GetTransactionService")
	defer span.End()

	resp := contract.GetTransactionResponse{}
	transaction := entity.Transaction{}
	err := app.GormDB().WithContext(ctx).Where("txn_id", txnID).Preload("Partner").Preload("Customer").First(&transaction).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	resp = contract.GetTransactionResponse{
		TxnID:             transaction.TxnID,
		OrderID:           transaction.OrderID,
		Phone:             transaction.Phone,
		AdminFee:          transaction.AdminFee,
		OTR:               transaction.OTR,
		InstallmentAmount: transaction.InstallmentAmount,
		InstallmentPeriod: transaction.InstallmentPeriod,
		Interest:          transaction.Interest,
		AssetName:         transaction.AssetName,
		TransactionDate:   transaction.TransactionDate,
		Status:            transaction.Status,
	}

	if transaction.Customer != nil {
		customer := transaction.Customer
		resp.Customer = &contract.GetTransactionCustomer{
			Email:     customer.Email,
			Phone:     customer.Phone,
			FullName:  customer.FullName,
			LegalName: customer.LegalName,
		}
	}

	if transaction.Partner != nil {
		partner := transaction.Partner
		resp.Partner = &contract.GetTransactionPartner{
			Name:    partner.Name,
			Address: partner.Address,
		}
	}

	return resp, nil
}

func (svc TransactionService) Settlement(ctx context.Context, req *contract.SettlementTransactionRequest) (contract.GetTransactionResponse, error) {
	ctx, span := app.Tracer().Start(ctx, "SettlementTransactionService")
	defer span.End()

	resp := contract.GetTransactionResponse{}

	db := app.GormDB().WithContext(ctx)

	// Get Customer Data
	customerID, err := middleware.GetAuthorizeCustomer(ctx)
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	customer := entity.Customer{}
	err = db.Preload("Limits").First(&customer, customerID).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	// Get Transaction
	transaction := entity.Transaction{}
	err = db.Where("txn_id", req.TxnID).First(&transaction).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	if transaction.Status != contract.TransactionPending && transaction.Status != contract.TransactionInit {
		err = errors.New("payment failed")
		tracer.RecordError(span, err)
		return resp, err
	}

	// Caculate possibility for the credit
	installmentPeriod := req.InstallmentPeriod
	otrWithAdminFee := transaction.OTR + transaction.AdminFee
	interestPercentage := customer.GetRiskPercentage()
	totalInterest := float64(otrWithAdminFee) * interestPercentage * float64(installmentPeriod) / 12
	total := float64(otrWithAdminFee) + totalInterest

	var customerLimit *entity.CustomerLimit
	for _, val := range customer.Limits {
		if val.Period == installmentPeriod {
			customerLimit = &val
			break
		}
	}

	if customerLimit == nil {
		err = fmt.Errorf("customer no support period %d", installmentPeriod)
		tracer.RecordError(span, err)
		return resp, err
	}

	totalInt := int(math.Ceil(total))
	unusedLimit := customerLimit.Limit - customerLimit.UsedLimit
	if unusedLimit < totalInt {
		err := errors.New("your limit isn't enough")
		tracer.RecordError(span, err)
		return resp, err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// Update Customer Limit
		for _, limit := range customer.Limits {
			if limit.Period <= installmentPeriod {
				limit.UsedLimit += totalInt

				err = db.Save(&limit).Error
				if err != nil {
					return err
				}
			}
		}

		// Update Transaction
		installment := math.Ceil(total / float64(installmentPeriod))

		transaction.Status = contract.TransactionSettlement
		transaction.InstallmentAmount = int(installment)
		transaction.InstallmentPeriod = installmentPeriod
		transaction.CustomerID = customerID
		transaction.Interest = float32(customer.Interest)

		err = db.Save(&transaction).Error
		if err != nil {
			return err
		}

		_, err = app.KafkaTransaction().WriteMessages(kafka.Message{
			Key:   []byte(contract.KafkaTransactionKey),
			Value: []byte(req.TxnID),
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	return svc.Get(ctx, req.TxnID)
}
