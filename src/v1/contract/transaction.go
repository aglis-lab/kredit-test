package contract

import "time"

// INIT, PENDING, SETTLEMENT, COMPLETE
const (
	TransactionInit       = "INIT"       // Partner initialize the transaction
	TransactionPending    = "PENDING"    // TBA,
	TransactionFailed     = "FAILED"     // TBA,
	TransactionCancelled  = "CANCELLED"  // User cancel or partner create new transaction from old order id
	TransactionSettlement = "SETTLEMENT" // User agree with the transaction or acknowledge
	TransactionComplete   = "COMPLETE"   // Kredit already paid partner
	TransactionFinish     = "FINISH"     // User already paid all the installments
)

type SettlementTransactionRequest struct {
	TxnID             string `json:"txn_id" validate:"required"`
	InstallmentPeriod int    `json:"installment_period" validate:"required"`
}

type TransactionRequest struct {
	APIKey  string `json:"api_key" validate:"required"`
	OrderID string `json:"order_id" validate:"required"`

	Phone     string `json:"phone" validate:"required"`
	OTR       int    `json:"otr" validate:"required"`
	AdminFee  int    `json:"admin_fee"`
	AssetName string `json:"asset_name" validate:"required"`
}

type GetTransactionPartner struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type GetTransactionCustomer struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FullName  string `json:"full_name"`
	LegalName string `json:"legal_name"`
}

type GetTransactionResponse struct {
	TxnID             string                  `json:"txn_id"`
	OrderID           string                  `json:"order_id"`
	Phone             string                  `json:"phone"`
	AdminFee          int                     `json:"admin_fee"`
	OTR               int                     `json:"otr"`
	InstallmentAmount int                     `json:"installment_amount"` // jumlah cicilan
	InstallmentPeriod int                     `json:"installment_period"`
	Interest          float32                 `json:"interest"` // In Percentage
	AssetName         string                  `json:"asset_name"`
	TransactionDate   time.Time               `json:"transaction_date"`
	Status            string                  `json:"status"`
	Partner           *GetTransactionPartner  `json:"partner,omitempty"`
	Customer          *GetTransactionCustomer `json:"customer,omitempty"`
}

type CreateTransactionResponse struct {
	TxnID   string `json:"txn_id"`
	OrderID string `json:"order_id"`
}

type CheckTransactionRequest struct {
	OTR               int `json:"otr" validate:"required"`
	InstallmentPeriod int `json:"installment_period" validate:"required"`
	AdminFee          int `json:"admin_fee" validate:"required"`
}

type InstallmentTransactionResponse struct {
}
