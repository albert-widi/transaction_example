package order

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/albert-widi/transaction_example/cmd/order/order/helper"
	"github.com/albert-widi/transaction_example/cmd/order/repo"
	"github.com/albert-widi/transaction_example/database"
	"github.com/albert-widi/transaction_example/log"
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
	ID               int64             `db:"id"`
	UserID           int64             `db:"user_id"`
	ShippingID       sql.NullInt64     `db:"shipping_id"`
	VoucherID        sql.NullInt64     `db:"voucher_id"`
	PaymentConfirmed sql.NullInt64     `db:"payment_confirmed"`
	Total            sql.NullInt64     `db:"total"`
	Status           OrderStatus       `db:"status"`
	CreatedAt        time.Time         `db:"created_at"`
	UpdatedAt        timeutil.NullTime `db:"updated_at"`
}

const (
	findActiveOrderByUserIDQuery = `
		SELECT id, user_id, shipping_id, voucher_id, payment_confirmed, status, total, created_at, updated_at
		FROM shop_order
		WHERE user_id = $1 AND status = $2
	`
)

func FindActiveOrderByUserID(userID int64) (Order, error) {
	o := Order{}
	err := database.MustGet(repo.DatabaseTX).Get(&o, findActiveOrderByUserIDQuery, userID, OrderStatusInit)
	if err != nil {
		log.Debug("[FindActiveOrderByUserID] Failed to get data: ", err.Error())
	}
	return o, err
}

// AddOrder struct
type AddOrderModel struct {
	UserID    int64
	ProductID int64 `json:"product_id"`
	Amount    int64 `json:"amount"`
}

func (a AddOrderModel) Validate() error {
	if a.UserID == 0 {
		return errors.New("UserID cannot be empty")
	}

	if a.ProductID == 0 {
		return errors.New("ProductID cannot be 0")
	}

	if a.Amount == 0 {
		return errors.New("Product amount cannot be 0")
	} else if a.Amount > 10 {
		return errors.New("Max amount of order is 10")
	}
	return nil
}

func AddOrder(ctx context.Context, add AddOrderModel) (int64, error) {
	err := add.Validate()
	if err != nil {
		return 0, err
	}

	// check product
	prod, err := helper.FindProductByID(ctx, add.ProductID)
	if err != nil {
		return 0, err
	}

	// check if an order is already created
	o, err := FindActiveOrderByUserID(add.UserID)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if o.ID == 0 {
		// create new order
		id, err := createOrder(add.UserID)
		if err != nil {
			return 0, err
		}
		o.ID = id
	}
	detail := OrderDetail{
		OrderID:   o.ID,
		ProductID: add.ProductID,
		Amount:    add.Amount,
		Price:     prod.Price,
		Total:     prod.Price * add.Amount,
	}
	err = createOrderDetail(detail)
	return o.ID, err
}

const (
	createOrderQuery = `INSERT INTO shop_order(user_id, status) VALUES($1,$2) RETURNING ID`
)

func createOrder(userID int64) (int64, error) {
	var orderID int64
	err := database.MustGet(repo.DatabaseTX).QueryRow(createOrderQuery, userID, OrderStatusInit).Scan(&orderID)
	return orderID, err
}

type OrderDetail struct {
	ID        int64 `db:"id"`
	OrderID   int64 `db:"order_id"`
	ProductID int64 `db:"product_id"`
	Amount    int64 `db:"amount"`
	Price     int64 `db:"price"`
	Total     int64 `db:"total"`
}

const (
	createOrderDetailQuery = `INSERT INTO shop_order_detail(order_id, product_id, amount, price, total) VALUES($1, $2, $3, $4, $5)`
)

func createOrderDetail(detail OrderDetail) error {
	_, err := database.MustGet(repo.DatabaseTX).Exec(createOrderDetailQuery, detail.OrderID, detail.ProductID, detail.Amount, detail.Price, detail.Total)
	return err
}
