package redis

import (
	"github.com/albert-widi/transaction_example/errors"
	"github.com/albert-widi/transaction_example/redis/internal"
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
		redisConn *redis_internal.Redis
		address   string
	}
	redis struct {
		connectedRedis map[redisType]*RedisStore
	}

	// type of redis
	redisType string
)

var redisObject *redis

// const of database type
const (
	SessionRedis redisType = "session_redis"
)

// Init redis connection
func Init(config Config) {
	redisObject = &redis{connectedRedis: make(map[redisType]*RedisStore)}
	for name, conf := range config.Redis {
		store := RedisStore{
			address: conf.Address,
		}
		store.redisConn = redis_internal.New(conf.Address, redis_internal.NetworkTCP, redis_internal.Options{
			MaxActive: conf.MaxActive,
			MaxIdle:   conf.MaxIdle,
			Timeout:   conf.Timeout,
			Wait:      true,
		})
		redisObject.connectedRedis[redisType(name)] = &store
	}
}

func Get(redisType redisType) (*RedisStore, error) {
	if redisConn, ok := redisObject.connectedRedis[redisType]; ok {
		return redisConn, nil
	}
	return nil, errors.New(errors.RedisNotExists)
}

func (r RedisStore) GetAddress() string {
	return r.address
}
