package order

import (
	"database/sql"
	"time"

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
	OrderID          int64             `json:"order_id,omitempty"`
	UserID           int64             `json:"user_id,omitempty"`
	ShippingID       sql.NullInt64     `json:"shipping_id,omitempty"`
	VoucherID        sql.NullInt64     `json:"voucher_id,omitempty"`
	PaymentConfirmed bool              `json:"payment_confirmed,omitempty"`
	Total            sql.NullInt64     `json:"total"`
	Status           OrderStatus       `json:"status,omitempty"`
	CreatedAt        time.Time         `json:"created_at,omitempty"`
	UpdatedAt        timeutil.NullTime `json:"updated_at"`
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

func AddOrder(userID, productID int64) error {
	// check if an order is already created
	o, err := FindActiveOrderByUserID(userID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if o.OrderID != 0 {
		// continue order
	} else {
		// create new order
	}
	// check product
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
