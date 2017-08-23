package apicalls

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	request "github.com/albert-widi/transaction_example/apicalls/internal"
	"github.com/albert-widi/transaction_example/log"
)

type authConfig struct {
	BaseURL string
}

type authAPI struct {
	Config authConfig
}

type authResp struct {
	Data interface{} `json:"data"`
}

func newAuth(config authConfig) *authAPI {
	api := &authAPI{
		Config: config,
	}
	return api
}

func (auth authAPI) Authenticate(cookie *http.Cookie, result interface{}) error {
	c, err := request.GetDefaultCleint()
	if err != nil {
		return err
	}
	req, err := http.NewRequest("GET", auth.Config.BaseURL+"/api/v1/auth", nil)
	if err != nil {
		return err
	}
	req.AddCookie(cookie)

	resp, err := c.DoRequest(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > http.StatusOK {
		return errors.New("Failed to authenticate user")
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Debugf("[Authenticate] Content: %s", string(content))

	authResp := &authResp{Data: result}
	err = json.Unmarshal(content, authResp)
	if err != nil {
		log.WithFields(log.Fields{"content": string(content)}).Debugf("[ERROR][GetProductByID] JsonUnmarshal: %s", err.Error())
	}
	if result != nil {
		log.Debugf("[Authenticate] Result: %+v", result)
	}
	return nil
}
