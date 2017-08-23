package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/albert-widi/transaction_example/cmd/order/order"
	"github.com/albert-widi/transaction_example/cmd/order/order/helper"
	"github.com/albert-widi/transaction_example/cmd/order/orderapi/middleware"
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
	// for test only
	w.Get("/ping", w.Handle(ping))
	w.Get("/test_auth", w.Handle(middleware.Authenticate(ping)))
	w.Get("/test", w.Handle(testProd))

	// all endpoints
	w.Post("/order", w.Handle(middleware.Authenticate(createNewOrder)))
}

func ping(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	resp.Data = map[string]string{
		"data": "This is data",
	}
	return resp, errors.DatabaseTypeNotExists.Err()
}

func testProd(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	p, err := helper.FindProductByID(r.Context(), 1)
	if err != nil {
		return resp, err
	}
	log.Debugf("Product: %+v", p)
	resp.Data = p
	return resp, nil
}

type newOrderRequest struct {
	ProductID int64 `json:"product_id"`
	Amount    int64 `json:"amount"`
}

func createNewOrder(r *http.Request) (route.HandleObject, error) {
	resp := new(route.V1)
	user, err := middleware.GetUser(r)
	if err != nil {
		return resp, err
	}

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return resp, errors.New(err, http.StatusBadRequest, "Failed to read order request")
	}

	req := order.AddOrderModel{}
	err = json.Unmarshal(content, &req)
	if err != nil {
		return resp, errors.New(err, http.StatusBadRequest, "Failed to read order request")
	}
	req.UserID = user.UserID

	orderID, err := order.AddOrder(r.Context(), req)
	if err != nil {
		return resp, err
	}
	resp.Data = map[string]interface{}{"order_id": orderID}
	return resp, nil
}
