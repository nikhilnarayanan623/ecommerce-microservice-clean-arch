package repository

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/domain"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/request"
	"github.com/nikhilnarayanan623/ecommerce-microservice-clean-arch/order-service/pkg/utils/response"
	"gorm.io/gorm"
)

type orderDatabase struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderDatabase{
		db: db,
	}
}

func (c *orderDatabase) Transaction(trxFun func(trxRepo interfaces.OrderRepository) error) error {
	trx := c.db.Begin()
	trxRepo := NewOrderRepository(trx)

	err := trxFun(trxRepo)
	if err != nil {
		trx.Rollback()
		return err
	}
	err = trx.Commit().Error
	if err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func (c *orderDatabase) SaveShopOrder(ctx context.Context, shopOrder domain.ShopOrder) (shopOrderID uint64, err error) {

	orderDate := time.Now()

	query := `INSERT INTO shop_orders (user_id, order_date, order_total_price, discount) VALUES($1, $2, $3, $4) RETURNING id`
	err = c.db.Raw(query, shopOrder.UserID, orderDate, shopOrder.OrderTotalPrice, shopOrder.Discount).Scan(&shopOrderID).Error

	return
}

func (c *orderDatabase) FindShopOrderByShopOrderID(ctx context.Context, shopOrderID uint64) (shopOrder domain.ShopOrder, err error) {

	query := `SELECT id, user_id, order_date, order_total_price, discount FROM shop_orders WHERE id = $1`
	err = c.db.Raw(query, shopOrderID).Scan(&shopOrder).Error

	return
}
func (c *orderDatabase) SaveOrderLine(ctx context.Context, orderLine domain.OrderLine) error {

	query := `INSERT INTO order_lines (product_item_id, shop_order_id, qty, price) VALUES ($1, $2, $3, $4)`
	err := c.db.Exec(query, orderLine.ProductItemID, orderLine.ShopOrderID, orderLine.Qty, orderLine.Price).Error

	return err
}
func (c *orderDatabase) FindOrderLinesByShopOrderID(ctx context.Context, shopOrderID uint64) (orderLines []domain.OrderLine, err error) {

	query := `SELECT id, product_item_id, shop_order_id, qty, price FROM order_lines WHERE shop_order_id = $1`
	err = c.db.Raw(query, shopOrderID).Scan(&orderLines).Error

	return
}

func (c *orderDatabase) FindAllShopOrdersByUserID(ctx context.Context, userID uint64, pagination request.Pagination) (shopOrders []response.ShopOrder, err error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT id, order_date, order_total_price, discount 
	FROM shop_orders WHERE user_id = $1
	LIMIT $2 OFFSET $3`
	err = c.db.Raw(query, userID, limit, offset).Scan(&shopOrders).Error

	return
}
