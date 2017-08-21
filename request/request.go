package request

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

//HTTPAPI for creating request or do request
type HTTPAPI struct {
	Method        string            // method
	URL           string            // url
	URIParams     map[string]string // params of API
	Body          []byte            // content body of API
	Headers       map[string]string // request header
	RestfulParams map[string]string // restful params for restful API
}

// Client struct
type Client struct {
	client         http.Client
	cachedTemplate map[string]*template.Template
	options        *Options

	mtx sync.RWMutex

	prometheusGaugeVec   *prometheus.GaugeVec
	prometheusSummaryVec *prometheus.SummaryVec
}

// Options for client
type Options struct {
	Timeout     int
	Environment string
	AutoLog     bool
	Monitor     Monitor
}

// Monitor struct to monitor request
type Monitor struct {
	Namespace string
	On        bool
}

const defaultTimeout int = 10

// New API object
func New(opt *Options) (*Client, error) {
	if opt == nil {
		opt = &Options{
			Timeout:     3,
			Environment: os.Getenv("TKPENV"),
			AutoLog:     false,
			Monitor:     Monitor{On: false},
		}
	}

	summaryVec := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: opt.Monitor.Namespace,
		Name:      "http_request_milisecond",
		Help:      "Average of http request response time at one time",
	}, []string{"endpoint", "method", "httpcode", "env"})
	if err := prometheus.Register(summaryVec); err != nil {
		// fmt.Printf("Failed to register prometheus metrics: %s", err.Error())
		return nil, fmt.Errorf("Failed to register prometheus metrics: %s", err.Error())
	}

	return &Client{
		client:               http.Client{Timeout: time.Second * time.Duration(opt.Timeout)},
		cachedTemplate:       make(map[string]*template.Template),
		prometheusSummaryVec: summaryVec,
		options:              opt,
	}, nil
}

func (c *Client) setTemplate(key string, t *template.Template) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.cachedTemplate[key] = t
}

func (c *Client) getTemplate(key string) *template.Template {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	return c.cachedTemplate[key]
}

// DoRequest for client http request
func (c *Client) DoRequest(request *http.Request) (*http.Response, error) {
	var (
		monitURL string
		code     string
	)

	defer func(timeStart time.Time) {
		if c.prometheusSummaryVec != nil {
			c.prometheusSummaryVec.With(prometheus.Labels{"endpoint": "all", "httpcode": code, "method": strings.ToUpper(request.Method), "env": c.options.Environment}).Observe(time.Since(timeStart).Seconds() * 1000)
			c.prometheusSummaryVec.With(prometheus.Labels{"endpoint": monitURL, "httpcode": code, "method": strings.ToUpper(request.Method), "env": c.options.Environment}).Observe(time.Since(timeStart).Seconds() * 1000)
		}
	}(time.Now())

	ctxVal := request.Context().Value("_monitoring_url")
	if ctxVal != nil {
		monitURL = ctxVal.(string)
	}

	resp, err := c.client.Do(request)
	// for monitoring
	if c.options.Monitor.On && monitURL != "" {
		code = "undefined"
		if resp != nil {
			code = fmt.Sprintf("%d", resp.StatusCode)
		} else if err != nil {
			code = "500"
			if err == context.Canceled || err == context.DeadlineExceeded {
				code = fmt.Sprintf("%d", http.StatusRequestTimeout)
			}
		}
	}
	return resp, err
}

// NewRequest for creating new http request
func (c *Client) NewRequest(h HTTPAPI) (*http.Request, error) {
	var (
		request          *http.Request
		buff             *bytes.Buffer
		rawURL, monitURL string
	)

	if h.URL == "" || len(h.URL) == 0 {
		return request, fmt.Errorf("URL is required")
	}

	h.Method = strings.ToUpper(h.Method)
	if h.Method != "POST" && h.Method != "GET" && h.Method != "PUT" && h.Method != "DELETE" {
		return request, fmt.Errorf("Unsupported method %s", h.Method)
	}

	if len(h.RestfulParams) > 0 {
		var (
			templateURL bytes.Buffer
			cached      bool
			t           *template.Template
		)
		// check if template already cached
		if temp := c.getTemplate(h.URL); temp != nil {
			t = temp
			cached = true
		} else {
			t = template.New(h.URL)
		}

		if err := t.Execute(&templateURL, h.RestfulParams); err != nil {
			return nil, err
		}
		monitURL = h.URL
		h.URL = templateURL.String()

		// cache template
		if !cached {
			c.setTemplate(h.URL, t)
		}
	} else {
		monitURL = h.URL
	}

	u, err := url.Parse(h.URL)
	if err != nil {
		return request, err
	}

	val := u.Query()
	for key, value := range h.URIParams {
		val.Add(key, value)
	}
	encodedVal := val.Encode()
	rawURL = u.String()

	if len(h.URIParams) > 0 {
		rawURL += "?" + encodedVal
	}

	switch h.Method {
	case "GET":
	case "POST":
		buff = bytes.NewBuffer(h.Body)
	case "PUT":
		buff = bytes.NewBuffer(h.Body)
	case "DELETE":
		buff = bytes.NewBuffer(h.Body)
	}

	//if buff == nil, will produce nil pointer interface
	if buff != nil {
		request, err = http.NewRequest(h.Method, rawURL, buff)
	} else {
		request, err = http.NewRequest(h.Method, rawURL, nil)
	}

	if err != nil {
		return request, err
	}

	for key, value := range h.Headers {
		request.Header.Add(key, value)
	}
	request.Header.Add("Content-Length", strconv.Itoa(len(encodedVal)))
	// add monit url to context
	request = request.WithContext(context.WithValue(request.Context(), "_monitoring_url", monitURL))
	return request, err
}
