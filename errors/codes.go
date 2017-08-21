package errors

import (
	"errors"
	"net/http"
)

// Codes of error
type Codes uint8

// error list
const (
	Other Codes = iota // Unclassified error.
	// data
	NoDataFound // no such data is exists

	// database error
	DatabaseTypeNotExists // Database type is not exists

	// redis
	RedisNotExists    // error because redis connection is not exists
	RedisKeyDuplicate // error duplicate because SetNX
	RedisKeyNotFound  // error because redis key not found

	// product
	ProductStockEmpty    // error product stock empty
	ProductInvalidAmount // invalid product amount

	// promo
	PromoVoucherNotExists // promo voucher is not exists
	PromoVoucherExpired   // promo voucher is expired
)

// Get code
func (c Codes) GetErrorAndCode() (string, int) {
	switch c {
	case Other:
		return "Internal server error", http.StatusInternalServerError
	case NoDataFound:
		return "No data found", http.StatusOK
	case DatabaseTypeNotExists:
		return "Database type is not exists", http.StatusInternalServerError
	case RedisNotExists:
		return "Redis is not exists", http.StatusInternalServerError
	case RedisKeyDuplicate:
		return "Redis key duplicate", http.StatusBadRequest
	case RedisKeyNotFound:
		return "Redis key is not found", http.StatusBadRequest
	case ProductStockEmpty:
		return "Product stock is empty", http.StatusBadRequest
	case ProductInvalidAmount:
		return "Invalid product amount", http.StatusBadRequest
	case PromoVoucherNotExists:
		return "Promo voucher not exists", http.StatusNotFound
	case PromoVoucherExpired:
		return "Promo voucher expired", http.StatusBadRequest
	default:
		return "Internal server error", http.StatusInternalServerError
	}
}

// Err return standard error type
func (c Codes) Err() error {
	errString, _ := c.GetErrorAndCode()
	return errors.New(errString)
}
