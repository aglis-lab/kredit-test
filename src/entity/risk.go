package entity

import "gorm.io/gorm"

type Risk struct {
	gorm.Model

	Interest float64 `db:"interest"`
	Risk     string  `db:"risk"`
	IsActive bool    `db:"is_active"`

	Limits []RiskLimit
}

func (risk *Risk) GetRiskPercentage() float64 {
	return risk.Interest / 100
}
