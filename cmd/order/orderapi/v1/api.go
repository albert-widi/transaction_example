package v1

import (
	"net/http"

	"github.com/albert-widi/transaction_example/cmd/order/order/helper"
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
	w.Get("/test", w.Handle(testProd))
}

func ping(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	resp.Data = map[string]string{
		"data": "This is data",
	}
	return resp, errors.DatabaseTypeNotExists.Err()
}

func testProd(r *http.Request) (route.HandleObject, error) {
	log.Debug("Going to testProd")
	resp := new(route.V1)
	p, err := helper.FindProductByID(r.Context(), 1)
	if err != nil {
		return resp, err
	}
	log.Debugf("Product: %+v", p)
	resp.Data = p
	return resp, nil
}
