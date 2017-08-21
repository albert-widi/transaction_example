package route

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/albert-widi/transaction_example/env"
	"github.com/albert-widi/transaction_example/log"
	"github.com/prometheus/client_golang/prometheus"
)

// HandleObject for handle
type HandleObject interface {
	// Render should handle the error themself and what to render afterwards
	// Even if return error is required, the error will be ignored by Handle middleware
	Render(http.ResponseWriter, *http.Request, error) error
}

// Handle type
type Handle func(*http.Request) (HandleObject, error)

// Handle custom http handler
func (rwrap *RouterWrapper) Handle(h Handle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tCurrent := time.Now()
		obj, err := h(r)

		// check if object is exists
		if obj == nil {
			log.Error("[FATAL] Handle return object is nil")
			http.Error(w, "Fatal Error", http.StatusInternalServerError)
			return
		}
		elapsedTime := time.Since(tCurrent).Seconds() * 1000
		ctxTime := context.WithValue(r.Context(), "_elapsed_time", elapsedTime)
		r = r.WithContext(ctxTime)
		obj.Render(w, r, err)
	}
}

// circuitbreaker
func circuitbreak(err error) bool {
	return false
}

func (rwrap *RouterWrapper) timeout(h http.HandlerFunc, options ...Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		opt := rwrap.opt
		if len(options) > 0 {
			opt = options[0]
		}
		// cancel context
		if opt.Timeout.Timeout > 0 {
			ctx, cancel := context.WithTimeout(r.Context(), time.Duration(opt.Timeout.Timeout)*time.Second)
			defer cancel()
			r = r.WithContext(ctx)
		}

		doneChan := make(chan bool)
		go func() {
			h(w, r)
			doneChan <- true
		}()
		select {
		case <-r.Context().Done():
			resp := opt.Timeout.Response
			if resp == nil {
				resp = defaultOptions.Timeout.Response
			}
			jsonResp, _ := json.Marshal(resp)
			w.WriteHeader(http.StatusRequestTimeout)
			w.Write(jsonResp)
			return
		case <-doneChan:
			return
		}
	}
}

func (rwrap *RouterWrapper) monitor(pattern string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := strings.ToUpper(r.Method)
		// send all metrics in defer
		defer func(timeStart time.Time) {
			var httpCode string
			code := r.Context().Value("HttpCode")
			if code == nil {
				httpCode = "599"
			} else {
				tmpCode := code.(int)
				httpCode = fmt.Sprintf("%d", tmpCode)
			}

			environment := env.Get()
			if prometheusSummaryVec != nil {
				prometheusSummaryVec.With(prometheus.Labels{"handler": "all", "method": method, "httpcode": httpCode, "env": environment}).Observe(time.Since(timeStart).Seconds() * 1000)
				prometheusSummaryVec.With(prometheus.Labels{"handler": pattern, "method": method, "httpcode": httpCode, "env": environment}).Observe(time.Since(timeStart).Seconds() * 1000)
			}
		}(time.Now())
		r = r.WithContext(context.WithValue(r.Context(), "HandlerFullPath", pattern))
		r = r.WithContext(context.WithValue(r.Context(), "HandlerMethod", method))
		h(w, r)
	}
}

// SetMonitorHttpCode will set http code to monitor
func (rwrap RouterWrapper) SetMonitorHttpCode(r *http.Request, httpcode int) {
	ctx := context.WithValue(r.Context(), "HttpCode", httpcode)
	r = r.WithContext(ctx)
}
