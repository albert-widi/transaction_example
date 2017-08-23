package shipping

import (
	"time"

	"github.com/albert-widi/transaction_example/cmd/logistic/repo"
	"github.com/albert-widi/transaction_example/cmd/logistic/shipper"
	"github.com/albert-widi/transaction_example/database"
	"github.com/albert-widi/transaction_example/errors"
	"github.com/albert-widi/transaction_example/timeutil"
)

type ShippingStatus int

const (
	ShippingInitiated  ShippingStatus = 10
	ShippingInProgress ShippingStatus = 20
	ShippingDelivered  ShippingStatus = 30
)

type Shipping struct {
	ID        int64             `db:"id"`
	ShipperID int64             `db:"shipper_id"`
	Price     int64             `db:"price"`
	Status    ShippingStatus    `db:"status"`
	From      string            `db:"from"`
	To        string            `db:"to"`
	CreatedAt time.Time         `db:"created_at"`
	UpdatedAt timeutil.NullTime `db:"updated_at"`
}

func (s Shipping) Validate() error {
	if s.ShipperID == 0 {
		return errors.New(errors.ShippingShipperEmpty)
	}
	if s.From == "" || s.To == "" {
		return errors.New(errors.ShippingAddressEmpty)
	}
	return nil
}

func (s Shipping) CanBeUpdated(status ShippingStatus) error {
	if s.Status > status {
		return errors.New(errors.ShippingCannotBeUpdated)
	}
	return nil
}

const (
	newShippingQuery = `INSERT INTO shipping(shipper_id, price, from, to, status) VALUES($1, $2, $3, $4, $5) RETURNING ID`
)

func NewShipping(shippingObject Shipping) (int64, error) {
	if err := shippingObject.Validate(); err != nil {
		return 0, err
	}
	sp, err := shipper.GetShipperByID(shippingObject.ShipperID)
	if err != nil {
		return 0, err
	}

	var shippingID int64
	err = database.MustGet(repo.DatabaseTX).
		QueryRow(newShippingQuery, sp.ID, sp.Price, shippingObject.From, shippingObject.To, ShippingInitiated).
		Scan(&shippingID)
	return shippingID, err
}

const (
	updateShippingStatusQuery = `UPDATE shipping SET status = $1 WHERE id = $2`
)

func UpdateShippingStatus(shippingID int64, status ShippingStatus) error {
	if shippingID == 0 {
		return errors.New(errors.ShippingIDInvalid)
	}
	s, err := FindShippingByID(shippingID)
	if err != nil {
		return err
	}
	// check if still can be updated
	if err := s.CanBeUpdated(status); err != nil {
		return err
	}

	_, err = database.MustGet(repo.DatabaseTX).
		Exec(updateShippingStatusQuery, shippingID, status)
	return err
}

const (
	findShippingByID = `SELECT id, shipper_id, price, status, from, to, created_at, updated_at
		FROM shipping
		WHERE id = $1`
)

func FindShippingByID(shippingID int64) (Shipping, error) {
	s := Shipping{}
	if shippingID == 0 {
		return s, errors.New(errors.ShippingIDInvalid)
	}
	err := database.MustGet(repo.DatabaseTX).
		Get(&s, findShippingByID, shippingID)
	return s, err
}
