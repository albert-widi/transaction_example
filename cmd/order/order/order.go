package order

import (
	"context"
	"database/sql"
	"time"

	"github.com/albert-widi/transaction_example/cmd/order/order/helper"
	"github.com/albert-widi/transaction_example/cmd/order/repo"
	"github.com/albert-widi/transaction_example/database"
	"github.com/albert-widi/transaction_example/timeutil"
)

type OrderStatus int

const (
	OrderStatusInit           OrderStatus = 10
	OrderStatusCancelled      OrderStatus = 20
	OrderStatusWaitingPayment OrderStatus = 25
	OrderStatusFinished       OrderStatus = 30
)

type Order struct {
	ID               int64             `database:"id"`
	UserID           int64             `database:"user_id"`
	ShippingID       sql.NullInt64     `database:"shipping_id"`
	VoucherID        sql.NullInt64     `database:"voucher_id"`
	PaymentConfirmed bool              `database:"payment_confirmed"`
	Total            sql.NullInt64     `database:"total"`
	Status           OrderStatus       `database:"status"`
	CreatedAt        time.Time         `database:"created_at"`
	UpdatedAt        timeutil.NullTime `database:"updated_at"`
}

const (
	findActiveOrderByUserIDQuery = `
		SELECT order_id, user_id, shipping_id, voucher_id, payment_confirmed, status, created_at, updated_at
		FROM shop_order
		WHERE user_id = $1 AND status = $2
	`
)

func FindActiveOrderByUserID(userID int64) (Order, error) {
	o := Order{}
	err := database.MustGet(repo.DatabaseTX).Get(&o, findActiveOrderByUserIDQuery, userID, OrderStatusInit)
	return o, err
}

func AddOrder(ctx context.Context, userID, productID int64) error {
	// check product
	prod, err := helper.FindProductByID(ctx, productID)
	if err != nil {
		return err
	}

	// check if an order is already created
	o, err := FindActiveOrderByUserID(userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if o.ID == 0 {
		// create new order
		id, err := createOrder(userID)
		if err != nil {
			return err
		}
		o.ID = id
	}
	err = createOrderDetail(o.ID, prod.ID, prod.Price)
	return err
}

const (
	createOrderQuery = `INSERT INTO shop_order(user_id, status) VALUES($1,$2) RETURNING ID`
)

func createOrder(userID int64) (int64, error) {
	var orderID int64
	err := database.MustGet(repo.DatabaseTX).QueryRow(createOrderQuery, userID, OrderStatusInit).Scan(&orderID)
	return orderID, err
}

const (
	createOrderDetailQuery = `INSERTINTO shop_order_detail(product_id, price) VALUES($1, $2, $3)`
)

func createOrderDetail(orderID, productID, productPrice int64) error {
	_, err := database.MustGet(repo.DatabaseTX).Exec(createOrderDetailQuery, orderID, productID, productPrice)
	return err
}

func ConfirmOrder(orderID int64) {

}
