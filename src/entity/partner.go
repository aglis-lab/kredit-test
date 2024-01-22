package entity

import "gorm.io/gorm"

type Partner struct {
	gorm.Model

	Name     string `db:"name"`
	Address  string `db:"address"`
	APIKey   string `db:"api_key"`
	IsActive bool   `db:"is_active"`
}
