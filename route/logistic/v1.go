package route_logistic

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/albert-widi/transaction_example/errors"
	"github.com/albert-widi/transaction_example/log"
)

// V1 struct
type V1 struct {
	Data    interface{}
	Message string
}

type v1Response struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   []string    `json:"errors,omitempty"`
	Status  string      `json:"status,omitempty"`
}

// Render function for V1
func (ws *V1) Render(w http.ResponseWriter, r *http.Request, err error) error {
	w.Header().Set("Content-Type", "application/json")
	var (
		e    []string
		code = http.StatusOK
	)
	status := "OK"
	if err != nil {
		var errMessage string
		status = "FAILED"
		errMessage, code = errors.ErrorAndHttpCode(err)
		e = append(e, errMessage)
		log.Error("Error when handling user request: ", err.Error())
	}

	response := v1Response{
		Data:    ws.Data,
		Message: ws.Message,
		Status:  status,
		Error:   e,
	}
	content, err := json.Marshal(&response)
	if err != nil {
		log.WithFields(log.Fields{"resp": fmt.Sprintf("%+v", response), "error": err.Error()}).
			Error("Failed to marshal JSON response.")
		resp := map[string]interface{}{
			"errors": []string{err.Error()},
			"status": "FAILED",
		}
		code = http.StatusInternalServerError
		content, _ = json.Marshal(resp)
	}
	w.WriteHeader(code)
	w.Write(content)
	return nil
}
