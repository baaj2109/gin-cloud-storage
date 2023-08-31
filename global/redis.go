package global

import (
	"time"

	"github.com/go-redis/redis"
)

type redisDB struct {
	db *redis.Client
}

var RedisDB *redisDB

func InitRedis() {
	RedisDB = NewRedis()
}

func NewRedis() *redisDB {
	// RedisDB = conn()
	return &redisDB{db: connToRedis()}
}

func connToRedis() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 10,
	})
	return rdb
}
func (r redisDB) GetKey(key string) (string, error) {
	return r.db.Get(key).Result()
}
func (r redisDB) SetKey(key string, value string, expires time.Duration) error {
	return r.db.Set(key, value, expires).Err()
}

func (r redisDB) DelKey(key string) error {
	return r.db.Del(key).Err()
}
