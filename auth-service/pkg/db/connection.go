package db

import (
	"fmt"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/auth-service/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		domain.RefreshSession{},
		domain.OtpSession{},
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}
