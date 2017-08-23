package voucher

import (
	"sync"
	"time"

	"github.com/albert-widi/transaction_example/errors"
)

// Product struct
type Voucher struct {
	ID              int64     `json:"id"`
	AmountRemaining int       `json:"amount_remaining"`
	Code            string    `json:"code"`
	Description     string    `json:"description"`
	Discount        int64     `json:"discount"`
	ExpireTime      time.Time `json:"expire_time"`
	mutex           sync.Mutex
}

// Hardcoded voucher list
var voucherList = map[int64]*Voucher{
	1: &Voucher{
		ID:              1,
		AmountRemaining: 5,
		Code:            "BELI1",
		Description:     "Beli 1 dapat 1",
		Discount:        10,
		ExpireTime:      time.Now().Add(time.Minute * 10),
	},
	2: &Voucher{
		ID:              2,
		AmountRemaining: 20,
		Code:            "BELICEPET",
		Description:     "Beli cepetan",
		Discount:        20,
		ExpireTime:      time.Now().Add(time.Second * 30),
	},
	3: &Voucher{
		ID:              3,
		AmountRemaining: 10,
		Code:            "BELIBESOK",
		Description:     "Beli 1 dapat 1",
		Discount:        5,
		ExpireTime:      time.Now().Add(time.Hour * 5),
	},
}

var voucherListBycode map[string]*Voucher

func init() {
	// re-map voucher list by code
	voucherListBycode = make(map[string]*Voucher)
	for _, value := range voucherList {
		voucherListBycode[value.Code] = value
	}
}

// pushing to use validation type for voucher
type voucherValid bool

const (
	Validate voucherValid = true
)

// Validate will validate voucher code for use
func ValidateVoucher(code string) error {
	v, ok := voucherListBycode[code]
	if !ok {
		return errors.New(errors.PromoVoucherNotExists)
	}
	v.mutex.Lock()
	defer v.mutex.Unlock()
	if time.Now().After(v.ExpireTime) {
		return errors.New(errors.PromoVoucherExpired)
	}
	v.AmountRemaining--
	return nil
}

// GetVoucherByCode return voucher and find by voucher code
// if validate param is passed, the function will validate the requirement for voucher
func GetVoucherByCode(code string, validate voucherValid) (*Voucher, error) {
	v, ok := voucherListBycode[code]
	if !ok {
		return v, errors.New(errors.PromoVoucherNotExists)
	}
	if validate {
		if time.Now().After(v.ExpireTime) {
			return v, errors.New(errors.PromoVoucherExpired)
		}
	}
	return v, nil
}
