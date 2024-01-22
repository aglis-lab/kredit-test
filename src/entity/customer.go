package entity

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model

	Email       string    `db:"email"`
	Phone       string    `db:"phone"`
	Password    string    `db:"password"`
	NIK         string    `db:"nik"`
	FullName    string    `db:"full_name"`
	LegalName   string    `db:"legal_name"`
	BornPlace   string    `db:"born_place"`
	BornAt      time.Time `db:"born_at"`
	Salary      int       `db:"salary"`
	PhotoKTP    string    `db:"photo_ktp"`
	PhotoSelfie string    `db:"photo_selfie"`

	Risk     string  `db:"risk"`
	Interest float64 `db:"interest"`

	Limits []CustomerLimit
}

func (customer *Customer) GetRiskPercentage() float64 {
	return customer.Interest / 100
}
