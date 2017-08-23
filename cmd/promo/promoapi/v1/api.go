package v1

import (
	"net/http"

	"github.com/albert-widi/transaction_example/cmd/promo/voucher"
	"github.com/albert-widi/transaction_example/errors"
	"github.com/albert-widi/transaction_example/log"
	"github.com/albert-widi/transaction_example/route"
	"github.com/pressly/chi"
)

// API struct
type APIV1 struct{}

// New api
func New() *APIV1 { return new(APIV1) }

// Register new api
func (api *APIV1) Register(r chi.Router) {
	log.Debug("Registering api v1 endpoints...")
	w := route.NewWrapper(r, route.Options{
		Timeout: route.Timeout{
			Timeout:  1,
			Response: map[string]string{"halo": "hola"},
		},
	})
	w.Get("/ping", w.Handle(ping))
	w.Get("/promo/voucher/validate", w.Handle(validateVoucher))
	w.Get("/promo/voucher/check", w.Handle(checkVoucher))
}

func ping(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	resp.Data = map[string]string{
		"data": "This is data",
	}
	return resp, errors.DatabaseTypeNotExists.Err()
}

func validateVoucher(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	code := r.FormValue("code")
	err := voucher.ValidateVoucher(code)
	if err != nil {
		return resp, errors.New(err)
	}
	resp.Message = "Voucher vaidated"
	return resp, nil
}

func checkVoucher(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	code := r.FormValue("code")
	log.Debug("[CheckVoucher] Code: ", code)
	v, err := voucher.GetVoucherByCode(code, voucher.Validate)
	if err != nil {
		return resp, errors.New(err)
	}
	resp.Data = v
	return resp, nil
}
