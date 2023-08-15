package repository

import (
	"context"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/cart-service/pkg/utils/response"
	"gorm.io/gorm"
)

type cartDatabase struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &cartDatabase{
		db: db,
	}
}

func (c *cartDatabase) FindCartByUserID(ctx context.Context, userID uint64) (cart domain.Cart, err error) {

	query := `SELECT id, user_id, total_price, applied_coupon_id, discount_amount FROM carts WHERE user_id = $1`
	err = c.db.Raw(query, userID).Scan(&cart).Error

	return
}
func (c *cartDatabase) SaveCart(ctx context.Context, userID uint64) (cartID uint64, err error) {

	query := `INSERT INTO carts (user_id, total_price) VALUES ($1, $2) RETURNING id`
	err = c.db.Raw(query, userID, 0).Scan(&cartID).Error

	return
}
func (c *cartDatabase) SaveCartItem(ctx context.Context, cartID, productItemID uint64) error {

	query := `INSERT INTO cart_items (cart_id, product_item_id, qty) VALUES ($1, $2, $3)`
	err := c.db.Exec(query, cartID, productItemID, 1).Error

	return err
}

func (c *cartDatabase) IsProductItemAlreadyExistInCart(ctx context.Context, cartID, productItemID uint64) (exist bool, err error) {

	query := `SELECT EXISTS(SELECT 1) AS exist FROM cart_items WHERE cart_id = $1 AND product_item_id = $2`
	err = c.db.Raw(query, cartID, productItemID).Scan(&exist).Error

	return
}

func (c *cartDatabase) FindCartItemsByCartID(ctx context.Context, cartID uint64) (cartItems []response.CartItem, err error) {

	query := `SELECT product_item_id, qty FROM cart_items WHERE cart_id = $1`
	err = c.db.Raw(query, cartID).Scan(&cartItems).Error

	return
}

func (c *cartDatabase) RemoveAllCartItems(ctx context.Context, cartID uint64) error {

	query := `DELETE FROM cart_items WHERE cart_id = $1`
	err := c.db.Exec(query, cartID).Error

	return err
}
