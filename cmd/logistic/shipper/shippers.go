package shipper

import (
	"github.com/albert-widi/transaction_example/errors"
)

// Shipper struct
type Shipper struct {
	ID    int64
	Name  string
	Price int64
}

func (sp Shipper) Validate() error {
	if sp.ID == 0 {
		return errors.New(errors.ShipperNotExists)
	}
	return nil
}

// let's assume that there is already a list of shippers
// and all the cost of shippers is fixed
var shippesList = map[int64]Shipper{
	1: Shipper{
		ID:    1,
		Name:  "Cepat",
		Price: 7,
	},
	2: Shipper{
		ID:    2,
		Name:  "Sedang",
		Price: 5,
	},
	3: Shipper{
		ID:    3,
		Name:  "Lambat",
		Price: 3,
	},
}

// GetShipperByID will find shipper based on ShipperID
func GetShipperByID(shipperID int64) (Shipper, error) {
	s, ok := shippesList[shipperID]
	if !ok {
		return s, errors.New(errors.ShipperNotExists)
	}
	return s, nil
}
