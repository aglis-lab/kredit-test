package entity

import (
	"gorm.io/gorm"
)

type CustomerLimit struct {
	gorm.Model

	CustomerID uint `db:"customer_id" gorm:"index"`
	Limit      int  `db:"limit"`
	Period     int  `db:"period"`
	UsedLimit  int  `db:"used_limit"`
}

func (limit *CustomerLimit) GetUsedLimit() int {
	if limit.Limit >= limit.UsedLimit {
		return limit.UsedLimit
	}

	return limit.Limit
}
