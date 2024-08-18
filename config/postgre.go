package config

import (
	"github.com/ThiagoKufa/estudo_go/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgre")

	dsn := "user=docker password=docker dbname=odontokufa sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorF("potgres opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("sqlite autoigration error: %v", err)
		return nil, err
	}

	return db, nil
}
