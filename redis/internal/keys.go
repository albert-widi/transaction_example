package redis_internal

import (
	redigo "github.com/garyburd/redigo/redis"
)

// Del key and value
func (r *Redis) Del(key string) (int64, error) {
	conn := r.Pool.Get()
	defer conn.Close()
	return redigo.Int64(conn.Do("DEL", key))
}
