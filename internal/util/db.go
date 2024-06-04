package util

import (
	"app/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	// dsn := "host=localhost user=yale password=12345 dbname=gowebapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(cfg.DB_URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
