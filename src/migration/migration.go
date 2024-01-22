package migration

import (
	"kreditplus/src/entity"

	"gorm.io/gorm"
)

func Init(gormDB *gorm.DB) error {
	err := gormDB.AutoMigrate(&entity.Customer{})
	if err != nil {
		return err
	}

	err = gormDB.AutoMigrate(&entity.CustomerLimit{})
	if err != nil {
		return err
	}

	err = gormDB.AutoMigrate(&entity.Transaction{})
	if err != nil {
		return err
	}

	err = gormDB.AutoMigrate(&entity.Partner{})
	if err != nil {
		return err
	}

	err = gormDB.AutoMigrate(&entity.Installment{})
	if err != nil {
		return err
	}

	err = gormDB.AutoMigrate(&entity.Risk{})
	if err != nil {
		return err
	}

	err = gormDB.AutoMigrate(&entity.RiskLimit{})
	if err != nil {
		return err
	}

	err = gormDB.AutoMigrate(&entity.RiskLimit{})
	if err != nil {
		return err
	}

	return nil
}
