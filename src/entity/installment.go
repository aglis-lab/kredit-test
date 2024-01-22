package entity

import (
	"time"

	"gorm.io/gorm"
)

type Installment struct {
	gorm.Model

	TransactionID     uint      `db:"transaction_id" gorm:"index"`
	InstallmentAmount int       `db:"installment_amount"`
	InstallmentPeriod int       `db:"installment_period"`
	InstallmentDate   time.Time `db:"installment_date"`
}
