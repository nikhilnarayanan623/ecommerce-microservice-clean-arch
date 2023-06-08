package db

import (
	"fmt"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/config"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		domain.Cart{},
		domain.CartItem{},
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	arrS := []byte(s)
	arrT := []byte(t)

	// set := map[byte]int{}

	// for _,char := range arrS{
	//     set[char]++
	// }

	// for _,char := range arrT{
	//     if count := set[char]; count == 0 {
	//         return false
	//     }
	//     set[char]--
	// }

	an := 0

	for i := 0; i < len(arrS); i++ {

		an ^= int(arrS[i])
		an ^= int(arrT[i])
	}

	return true
}
