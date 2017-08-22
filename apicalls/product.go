package apicalls

type productConfig struct {
	BaseURL string
}

type ProductAPI struct {
	Config productConfig
}

func newProduct(config productConfig) *ProductAPI {
	api := &ProductAPI{
		Config: config,
	}
	return api
}
