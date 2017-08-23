package apicalls

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	request "github.com/albert-widi/transaction_example/apicalls/internal"
	"github.com/albert-widi/transaction_example/errors"
	"github.com/albert-widi/transaction_example/log"
)

type productConfig struct {
	BaseURL string
}

type productAPI struct {
	Config productConfig
}

type prodResponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

func newProduct(config productConfig) *productAPI {
	api := &productAPI{
		Config: config,
	}
	return api
}

func (prod productAPI) GetProductByID(ctx context.Context, productID int64, result interface{}) error {
	url := "/api/v1/product/" + fmt.Sprintf("%d", productID)
	resp, err := request.DoRequestWithContext(ctx, request.HTTPAPI{
		Method: "GET",
		URL:    prod.Config.BaseURL + url,
	})
	if err != nil {
		log.Debugf("[ERROR][GetProductByID] DoRequestWithContext: %s", err.Error())
		return err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Debugf("[GetProductByID] Content: %s", string(content))

	prodResp := &prodResponse{Data: result}
	err = json.Unmarshal(content, prodResp)
	if err != nil {
		log.WithFields(log.Fields{"content": string(content)}).Debugf("[ERROR][GetProductByID] JsonUnmarshal: %s", err.Error())
	}
	if result != nil {
		log.Debugf("[GetProductByID] Result: %+v", result)
	}
	return err
}

func (prod productAPI) DecreaseProductStock(ctx context.Context, productID, amount int64) error {
	url := prod.Config.BaseURL + "/api/v1/product/" + fmt.Sprintf("%d", productID) + "/decrease"
	log.Debug("[DecreaseProductStock] URL: ", url)
	resp, err := request.DoRequestWithContext(ctx, request.HTTPAPI{
		Method:    "PUT",
		URL:       url,
		URIParams: map[string]string{"amount": fmt.Sprintf("%d", amount)},
	})
	if err != nil {
		log.Debugf("[ERROR][DecreaseProductStock] DoRequestWithContext: %s", err.Error())
		return err
	}
	defer resp.Body.Close()

	// for debugging purpose only
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Debug("[DecreaseProductStock] Content: ", string(content))

	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to decrease product quantity")
	}
	return err
}
