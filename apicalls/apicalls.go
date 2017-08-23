package apicalls

import (
	request "github.com/albert-widi/transaction_example/apicalls/internal"
	"github.com/albert-widi/transaction_example/env"
)

// Config of apicalls
type Config struct {
	Auth     authConfig
	Logistic logsiticConfig
	Product  productConfig
	Promo    promoConfig
}

// api list
var (
	Auth     *authAPI
	Logistic *LogisticAPI
	Product  *productAPI
	Promo    *PromoAPI
)

// New apicalls
func New(conf Config) error {
	appName, err := env.GetAppName()
	if err != nil {
		return err
	}
	// need to init request for request default client
	if err := request.Init(request.Options{
		Environment: env.Get(),
		Timeout:     3,
		Monitor:     request.Monitor{Namespace: appName, On: true},
	}); err != nil {
		return err
	}
	// list of available API
	Auth = newAuth(conf.Auth)
	Logistic = newLogistic(conf.Logistic)
	Product = newProduct(conf.Product)
	Promo = newPromo(conf.Promo)
	return nil
}
