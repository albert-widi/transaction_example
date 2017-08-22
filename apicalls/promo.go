package apicalls

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
