package apiv1

import (
	"net/http"

	"github.com/albert-widi/transaction_example/errors"
	"github.com/albert-widi/transaction_example/route"
	"github.com/alileza/common/log"
	"github.com/pressly/chi"
)

// API struct
type API struct{}

// New api
func New() *API { return new(API) }

// Register new api
func (api *API) Register(r chi.Router) {
	log.Debug("Registering apiv1 endpoints...")
	w := route.NewWrapper(r, route.Options{
		Timeout: route.Timeout{
			Timeout:  1,
			Response: map[string]string{"halo": "hola"},
		},
	})
	w.Get("/ping", w.Handle(ping))
}

func ping(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	resp.Data = map[string]string{
		"data": "This is data",
	}
	return resp, errors.DatabaseTypeNotExists.Err()
}
