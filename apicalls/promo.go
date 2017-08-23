package apicalls

import (
	"context"
	"errors"
	"net/http"

	request "github.com/albert-widi/transaction_example/apicalls/internal"
)

type promoConfig struct {
	BaseURL string
}

type PromoAPI struct {
	Config promoConfig
}

func newPromo(config promoConfig) *PromoAPI {
	api := &PromoAPI{
		Config: config,
	}
	return api
}

func (promo PromoAPI) ValidateVoucherByCode(ctx context.Context, voucherCode string) error {
	resp, err := request.DoRequestWithContext(ctx, request.HTTPAPI{
		Method:    "GET",
		URL:       promo.Config.BaseURL + "/api/v1/promo/voucher/validate",
		URIParams: map[string]string{"code": voucherCode},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to validate voucher code")
	}
	return nil
}
