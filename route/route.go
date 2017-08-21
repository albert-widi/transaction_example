package route

import (
	"net/http"

	"github.com/albert-widi/transaction_example/log"
	"github.com/pressly/chi"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	prometheusSummaryVec *prometheus.SummaryVec
	defaultOptions       = Options{
		Timeout: Timeout{
			Timeout: 2,
			Response: map[string]string{
				"Status": "Request time out",
			},
		},
		Circuitbreak: Circuitbreak{
			Breaker: true,
			Response: map[string]string{
				"Status": "Service unavailable",
			},
		},
	}
)

// Options for timeout and debugging
type Options struct {
	Pattern      string // for grouping
	Timeout      Timeout
	Circuitbreak Circuitbreak
}

// Timeout struct for response
type Timeout struct {
	Timeout  int
	Response interface{}
}

// Circuitbreak for circuitbreaker options
type Circuitbreak struct {
	Breaker  bool
	Response interface{}
}

// RouterWrapper wrap chi route
type RouterWrapper struct {
	r   chi.Router
	opt Options
}

func init() {
	prometheusSummaryVec = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "transactionapp",
		Name:      "handler_request_milisecond",
		Help:      "Average of handler response time at one time",
	}, []string{"handler", "method", "httpcode", "env"})
	if err := prometheus.Register(prometheusSummaryVec); err != nil {
		log.Errorf("Failed to register prometheus metrics: %s", err.Error())
	}
}

// NewWrapper to wrap route
func NewWrapper(router chi.Router, options ...Options) *RouterWrapper {
	var o Options
	if len(options) > 0 {
		o = options[0]
	}
	return &RouterWrapper{
		r:   router,
		opt: o,
	}
}

// Get router
func (rwrap *RouterWrapper) Get(pattern string, h http.HandlerFunc, options ...Options) {
	log.Debugf("Initializing GET for %s", pattern)
	rwrap.r.Get(rwrap.opt.Pattern+pattern, rwrap.monitor(pattern, rwrap.timeout(h, options...)))
}

// Post router
func (rwrap *RouterWrapper) Post(pattern string, h http.HandlerFunc, options ...Options) {
	log.Debugf("Initializing POST for %s", pattern)
	rwrap.r.Post(rwrap.opt.Pattern+pattern, rwrap.monitor(pattern, rwrap.timeout(h, options...)))
}

// Put router
func (rwrap *RouterWrapper) Put(pattern string, h http.HandlerFunc, options ...Options) {
	log.Debugf("Initializing PUT for %s", pattern)
	rwrap.r.Put(rwrap.opt.Pattern+pattern, rwrap.monitor(pattern, rwrap.timeout(h, options...)))
}

// Delete router
func (rwrap *RouterWrapper) Delete(pattern string, h http.HandlerFunc, options ...Options) {
	log.Debugf("Initializing DELETE for %s", pattern)
	rwrap.r.Delete(rwrap.opt.Pattern+pattern, rwrap.monitor(pattern, rwrap.timeout(h, options...)))
}

// HandlePrometheus for handling prometheus metrics
func (rwrap *RouterWrapper) HandlePrometheus(pattern string, h http.Handler) {
	rwrap.r.Handle(pattern, h)
}

// Param return string value from chi.URLParam
func Param(r *http.Request, param string) string {
	return chi.URLParam(r, param)
}
