package request

import (
	"context"
	"errors"
	"net/http"
)

var (
	ErrDefaultClientIsNil = errors.New("Default client is nil")
)

//Params for HTTPAPI
type Params map[string]string

//HeaderOptions for HTTPAPI
type HeaderOptions map[string]string

//DefaultClient only need one client, since this client will spawn goroutines in every request
var defaultClient *Client

func Init(opt Options) error {
	var err error
	defaultClient, err = New(&opt)
	return err
}

// NewRequest for creating new http request. Will deprecate this function soon
func NewRequest(h HTTPAPI) (*http.Request, error) {
	if defaultClient == nil {
		return nil, ErrDefaultClientIsNil
	}
	return defaultClient.NewRequest(h)
}

func doRequest(h HTTPAPI, ctxs ...context.Context) (*http.Response, error) {
	if defaultClient == nil {
		return nil, ErrDefaultClientIsNil
	}
	var response *http.Response

	request, err := defaultClient.NewRequest(h)
	if err != nil {
		return response, err
	}

	for _, val := range ctxs {
		request = request.WithContext(val)
	}
	response, err = defaultClient.DoRequest(request)
	return response, err
}

//DoRequest will do http client request
func DoRequest(h HTTPAPI) (*http.Response, error) {
	return doRequest(h)
}

// DoRequestWithContext will request data using provided context
func DoRequestWithContext(ctx context.Context, h HTTPAPI) (*http.Response, error) {
	return doRequest(h, ctx)
}

// DoRequestWithMultipleContext allowing to pass multiple context to the request
func DoRequestWithMultipleContext(ctxs []context.Context, h HTTPAPI) (*http.Response, error) {
	return doRequest(h, ctxs...)
}
