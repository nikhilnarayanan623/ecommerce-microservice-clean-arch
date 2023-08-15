package db

import (
	"fmt"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/product-service/pkg/domain"
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
		domain.Category{},
		domain.Variation{}, domain.VariationOption{},
		domain.Product{}, domain.ProductItem{},
		domain.ProductConfiguration{},
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}
