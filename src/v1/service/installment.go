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

	"gorm.io/gorm"
)

func (svc InstallmentService) Calculation(ctx context.Context, req *contract.CalculationInstallmentRequest) (contract.CalculationInstallmentResponse, error) {
	ctx, span := app.Tracer().Start(ctx, "CalculationInstallmentService")
	defer span.End()

	resp := contract.CalculationInstallmentResponse{}

	db := app.GormDB().WithContext(ctx)

	// Fetch customer limit
	risk := entity.Risk{}
	err := db.Model(&entity.Risk{}).Preload("Limits").Where("risk", contract.DefaultRisk).Find(&risk).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	interestPercentage := .0
	if customerID, err := middleware.GetAuthorizeCustomer(ctx); err == nil {
		customer := entity.Customer{}

		err = db.First(&customer, customerID).Error
		if err != nil {
			tracer.RecordError(span, err)
			return resp, err
		}

		interestPercentage = customer.GetRiskPercentage()
		resp.Interest = customer.Interest
		resp.Risk = customer.Risk
	} else {
		interestPercentage = risk.GetRiskPercentage()
		resp.Interest = risk.Interest
		resp.Risk = risk.Risk
	}

	// Calculate Installment based on limits
	for _, limit := range risk.Limits {
		totalInterest := float64(req.OTR) * interestPercentage * float64(limit.Period) / 12
		installment := (int64(math.Ceil(totalInterest)) + req.OTR) / int64(limit.Period)

		resp.Installments = append(resp.Installments, contract.CalculatiomInstallmentItem{
			Installment: installment,
			Period:      limit.Period,
			IsAllowed:   true,
		})
	}

	resp.OTR = req.OTR

	return resp, nil
}

func (svc InstallmentService) CalculationTransaction(ctx context.Context, txnID string) (contract.CalculationInstallmentResponse, error) {
	ctx, span := app.Tracer().Start(ctx, "CalculationTransactionInstallmentService")
	defer span.End()

	resp := contract.CalculationInstallmentResponse{}

	db := app.GormDB().WithContext(ctx)

	// Fetch customer limit
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

	// Transaction
	transaction := entity.Transaction{}
	err = db.Where("txn_id", txnID).First(&transaction).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	// Calculate Installment based on limits
	otr := transaction.OTR + transaction.AdminFee
	interestPercentage := customer.GetRiskPercentage()
	for _, limit := range customer.Limits {
		totalInterest := float64(otr) * interestPercentage * float64(limit.Period) / 12
		installment := (int64(math.Ceil(totalInterest)) + int64(otr)) / int64(limit.Period)

		resp.Installments = append(resp.Installments, contract.CalculatiomInstallmentItem{
			Installment: installment,
			Period:      limit.Period,
			IsAllowed:   (limit.Limit - limit.UsedLimit) > otr,
		})
	}

	resp.OTR = int64(otr)
	resp.Interest = customer.Interest
	resp.Risk = customer.Risk

	return resp, nil
}

func (svc InstallmentService) Payment(ctx context.Context, req *contract.PaymentInstallmentRequest) (contract.PaymentInstallmentResponse, error) {
	ctx, span := app.Tracer().Start(ctx, "PaymentInstallmentService")
	defer span.End()

	resp := contract.PaymentInstallmentResponse{}

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

	if transaction.Status != contract.TransactionComplete {
		err := errors.New("installment can't be made")
		tracer.RecordError(span, err)
		return resp, err
	}

	// Check if payment for those period already exist
	count := int64(0)
	err = db.Model(&entity.Installment{}).Where("transaction_id", transaction.ID).Where("installment_period", req.InstallmentPeriod).Count(&count).Error
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	if count > 0 {
		err := fmt.Errorf("installment for period %d already paid", req.InstallmentPeriod)
		tracer.RecordError(span, err)
		return resp, err
	}

	// Doing Transaction
	installmentDate := time.Now()
	err = db.Transaction(func(tx *gorm.DB) error {
		// Create Installment
		installment := entity.Installment{
			TransactionID:     transaction.ID,
			InstallmentAmount: transaction.InstallmentAmount,
			InstallmentPeriod: req.InstallmentPeriod,
			InstallmentDate:   installmentDate,
		}
		err = db.Create(&installment).Error
		if err != nil {
			return err
		}

		// Reduce Limit
		for _, limit := range customer.Limits {
			if limit.Period <= transaction.InstallmentPeriod {
				limit.UsedLimit -= transaction.InstallmentAmount

				err = db.Save(&limit).Error
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		tracer.RecordError(span, err)
		return resp, err
	}

	resp = contract.PaymentInstallmentResponse{
		TransactionID:     transaction.ID,
		InstallmentAmount: transaction.InstallmentAmount,
		InstallmentPeriod: req.InstallmentPeriod,
		InstallmentDate:   installmentDate,
	}

	return resp, nil
}
