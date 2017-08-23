package order

import (
	"context"
	"time"

	"github.com/albert-widi/transaction_example/cmd/order/order/helper"
	"github.com/albert-widi/transaction_example/cmd/order/repo"
	"github.com/albert-widi/transaction_example/database"
	"github.com/albert-widi/transaction_example/errors"
	"github.com/albert-widi/transaction_example/log"
	"github.com/albert-widi/transaction_example/timeutil"
)

type SubmitOrderModel struct {
	OrderID     int64
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	VoucherCOde string `json:"voucher_code"`
}

func (sbm SubmitOrderModel) Validate() error {
	if sbm.OrderID == 0 {
		return errors.New("OrderID cannot be empty")
	}

	if sbm.Name == "" || sbm.Phone == "" || sbm.Email == "" || sbm.Address == "" {
		return errors.New("Name, phone number, email, address cannot be empty")
	}
	return nil
}

func SubmitOrder(ctx context.Context, model SubmitOrderModel) error {
	if err := model.Validate(); err != nil {
		return err
	}

	o, err := FindOrderByID(model.OrderID)
	if err != nil {
		log.Debug("[SubmitOrder] Failed find order by ID: ", err.Error())
		return err
	}
	if err := o.IsSubmitable(); err != nil {
		return err
	}

	err = validateProductAndVoucher(o.ID, model.VoucherCOde)
	if err != nil {
		log.Debug("[SubmitOrder] Failed to validate product and voucher: ", err.Error())
		return err
	}

	err = o.UpdateStatus(OrderStatusWaitingPayment)
	if err != nil {
		log.Debugf("[SubmitOrder] Failed to update status: ", err.Error())
		return err
	}

	// submit customer detail
	customerDetail := OrderCustomerDetail{
		OrderID:     model.OrderID,
		Name:        model.Name,
		PhoneNumber: model.Phone,
		Email:       model.Email,
		Address:     model.Address,
	}
	err = CreateOrderCustomerDetail(customerDetail)
	if err != nil {

	}
	return err
}

func validateProductAndVoucher(orderID int64, voucherCode string) error {
	details, err := FindOrderDetailByOrderID(orderID)
	if err != nil {
		return err
	}
	for _, od := range details {
		err := helper.DecreaseProductStock(context.Background(), od.ProductID, od.Amount)
		if err != nil {
			return err
		}
		err = helper.ValidateVoucher(context.Background(), voucherCode)
		if err != nil {
			return err
		}
	}
	return nil
}

type OrderCustomerDetail struct {
	ID          int64             `db:"id"`
	OrderID     int64             `db:"order_id"`
	Name        string            `db:"name"`
	PhoneNumber string            `db:"phone_number"`
	Email       string            `db:"email"`
	Address     string            `db:"address"`
	CreatedAt   time.Time         `db:"crated_at"`
	UpdatedAt   timeutil.NullTime `db:"updated_at"`
}

const (
	createOrderCustomerDetailQuery = `
		INSERT INTO order_customer_detail(order_id, name, phone_number, email, address)
		VALUES($1, $2, $3, $4, $5)
	`
)

func CreateOrderCustomerDetail(model OrderCustomerDetail) error {
	_, err := database.MustGet(repo.DatabaseTX).Exec(createOrderCustomerDetailQuery, model.OrderID, model.Name, model.PhoneNumber, model.Email, model.Address)
	return err
}
