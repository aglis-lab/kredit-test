package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model

	CustomerID uint   `db:"customer_id" gorm:"index"`
	PartnerID  uint   `db:"partner_id" gorm:"index"`
	OrderID    string `db:"order_id" gorm:"index"`

	Phone             string    `db:"phone"`
	AdminFee          int       `db:"admin_fee"`
	OTR               int       `db:"otr"`
	InstallmentAmount int       `db:"installment_amount"` // jumlah cicilan
	InstallmentPeriod int       `db:"installment_period"`
	Interest          float32   `db:"interest"`
	AssetName         string    `db:"asset_name"`
	TransactionDate   time.Time `db:"transaction_date"`

	TxnID  string `db:"txn_id" gorm:"index"`
	Status string `db:"status"`

	Customer *Customer
	Partner  *Partner
}
