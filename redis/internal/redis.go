package redis

import (
	"sync"
	"time"

	redigo "github.com/garyburd/redigo/redis"
)

// NetworkTCP for tcp
const NetworkTCP = "tcp"

// Options for redis
type Options struct {
	MaxIdle   int
	MaxActive int
	Timeout   int
	Wait      bool
}

// Redis struct
type Redis struct {
	Pool  *redigo.Pool
	mutex sync.Mutex
}

// New redis connection
func New(address, network string, opts ...Options) *Redis {
	opt := Options{
		MaxIdle:   100,
		MaxActive: 100,
		Timeout:   100,
		Wait:      true,
	}
	if len(opts) > 0 {
		opt = opts[0]
	}

	return &Redis{
		Pool: &redigo.Pool{
			MaxIdle:     opt.MaxIdle,
			MaxActive:   opt.MaxActive,
			IdleTimeout: time.Duration(opt.Timeout) * time.Second,
			Dial: func() (redigo.Conn, error) {
				return redigo.Dial(network, address)
			},
		},
	}
}
