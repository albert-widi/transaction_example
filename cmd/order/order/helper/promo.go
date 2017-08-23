package helper

import (
	"context"

	"github.com/albert-widi/transaction_example/apicalls"
)

func ValidateVoucher(ctx context.Context, voucherCode string) error {
	if voucherCode == "" {
		return nil
	}
	return apicalls.Promo.ValidateVoucherByCode(ctx, voucherCode)
}
