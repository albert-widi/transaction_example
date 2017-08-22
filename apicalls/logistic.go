package apicalls

type logsiticConfig struct {
	BaseURL string
}

type LogisticAPI struct {
	Config logsiticConfig
}

func newLogistic(config logsiticConfig) *LogisticAPI {
	api := &LogisticAPI{
		Config: config,
	}
	return api
}
