package entity

import (
	"gorm.io/gorm"
)

type RiskLimit struct {
	gorm.Model

	RiskID uint `db:"risk_id" gorm:"index"`
	Limit  int  `db:"limit"`
	Period int  `db:"period"`
}
