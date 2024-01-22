package contract

import "time"

type PaymentInstallmentRequest struct {
	TxnID             string `json:"txn_id" validate:"required"`
	InstallmentPeriod int    `json:"installment_period" validate:"required"`
}

type CalculationInstallmentRequest struct {
	OTR int64 `json:"otr" validate:"required"`
}

type PaymentInstallmentResponse struct {
	TransactionID     uint      `json:"transaction_id"`
	InstallmentAmount int       `json:"installment_amount"`
	InstallmentPeriod int       `json:"installment_period"`
	InstallmentDate   time.Time `json:"installment_date"`
}

type CalculationInstallmentResponse struct {
	OTR          int64                        `json:"otr"`
	Risk         string                       `json:"risk"`
	Interest     float64                      `json:"interest"`
	Installments []CalculatiomInstallmentItem `json:"installments"`
}

type CalculatiomInstallmentItem struct {
	Installment int64 `json:"installment"`
	Period      int   `json:"period"`
	IsAllowed   bool  `json:"is_allowed"`
}
