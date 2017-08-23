package apicalls

type authConfig struct {
	BaseURL string
}

type authAPI struct {
	Config authConfig
}

func newAuth(config authConfig) *authAPI {
	api := &authAPI{
		Config: config,
	}
	return api
}
