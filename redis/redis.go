package redis

import (
	"github.com/albert-widi/transaction_example/errors"
	redi "github.com/albert-widi/transaction_example/redis/internal"
)

type (
	redisconf struct {
		Address   string
		MaxActive int
		MaxIdle   int
		Timeout   int
	}
	Config struct {
		Redis map[string]*redisconf
	}

	RedisStore struct {
		redisConn *redi.Redis
		address   string
	}
	redis struct {
		connectedRedis map[RedisType]*RedisStore
	}

	// type of redis
	RedisType string
)

var redisObject *redis

// Init redis connection
func Init(config Config) {
	redisObject = &redis{connectedRedis: make(map[RedisType]*RedisStore)}
	for name, conf := range config.Redis {
		store := RedisStore{
			address: conf.Address,
		}
		store.redisConn = redi.New(conf.Address, redi.NetworkTCP, redi.Options{
			MaxActive: conf.MaxActive,
			MaxIdle:   conf.MaxIdle,
			Timeout:   conf.Timeout,
			Wait:      true,
		})
		redisObject.connectedRedis[RedisType(name)] = &store
	}
}

func Get(redisType RedisType) (*RedisStore, error) {
	if redisConn, ok := redisObject.connectedRedis[redisType]; ok {
		return redisConn, nil
	}
	return nil, errors.New(errors.RedisNotExists)
}

func (r RedisStore) GetAddress() string {
	return r.address
}
