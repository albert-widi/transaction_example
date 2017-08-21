package redis

import (
	"github.com/albert-widi/transaction_example/errors"

	redigo "github.com/garyburd/redigo/redis"
)

// SetNotExists function
func (r *RedisStore) SetNotExists(key, value string) error {
	_, err := r.redisConn.SetWithNX(key, value, 10)
	if err != nil {
		if err == redigo.ErrNil {
			return errors.New(errors.RedisKeyDuplicate)
		}
		return err
	}
	return nil
}

// Set function
func (r *RedisStore) Set(key, value string, expire int) error {
	_, err := r.redisConn.SetEX(key, value, expire)
	return err
}

// Get function
func (r *RedisStore) Get(key string) (string, error) {
	value, err := r.redisConn.Get(key)
	if err == redigo.ErrNil {
		return "", errors.New(errors.RedisKeyNotFound)
	}
	return value, err
}

// Del function
func (r *RedisStore) Del(key string) error {
	_, err := r.redisConn.Del(key)
	if err == redigo.ErrNil {
		return errors.New(errors.RedisKeyNotFound)
	}
	return err
}
