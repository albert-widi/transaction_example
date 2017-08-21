package v1

import (
	"net/http"
	"strconv"

	"github.com/albert-widi/transaction_example/cmd/product/product"
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
	w.Get("/products", w.Handle(getProductList))
	w.Put("/product/{id}/increase", w.Handle(increaseStock))
	w.Put("/product/{id}/decrease", w.Handle(decreaseStock))
}

func ping(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	resp.Data = map[string]string{
		"data": "This is data",
	}
	return resp, errors.DatabaseTypeNotExists.Err()
}

func getProductList(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	list, err := product.GetProductList()
	if err != nil {
		resp.Message = "Failed to get product list"
		return resp, errors.New(err)
	}
	resp.Data = list
	return resp, nil
}

func productAmountFromReq(r *http.Request) (int, error) {
	amount, err := strconv.Atoi(r.FormValue("amount"))
	if err != nil {
		return 0, errors.New(err)
	}
	if amount <= 0 {
		return 0, errors.New(errors.ProductInvalidAmount)
	}
	return amount, nil
}

func increaseStock(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	productID, err := strconv.ParseInt(route.Param(r, "id"), 10, 64)
	if err != nil {
		resp.Message = "Invalid product_id"
		return resp, errors.New(err)
	}
	amount, err := productAmountFromReq(r)
	if err != nil {
		return resp, err
	}

	err = product.IncreaseStock(productID, amount)
	if err != nil {
		resp.Message = "Failed to increase product stock"
		return resp, errors.New(err)
	}
	resp.Message = "Success increase product stock"
	return resp, nil
}

func decreaseStock(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	productID, err := strconv.ParseInt(route.Param(r, "id"), 10, 64)
	if err != nil {
		resp.Message = "Invalid product_id"
		return resp, errors.New(err)
	}
	amount, err := productAmountFromReq(r)
	if err != nil {
		return resp, err
	}

	err = product.DecreaseStock(productID, amount)
	if err != nil {
		resp.Message = "Failed to decrease product stock"
		return resp, errors.New(err)
	}
	resp.Message = "Success decrease product stock"
	return resp, nil
}
