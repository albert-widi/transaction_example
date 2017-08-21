package productapi

import (
	apiv1 "github.com/albert-widi/transaction_example/cmd/product/productapi/v1"
	"github.com/pressly/chi"
)

// Endpoints describe all grouped endpoints in the application
var Endpoints map[string]func(chi.Router) = map[string]func(chi.Router){
	"/api/v1": func(r chi.Router) { apiv1.New().Register(r) },
}
