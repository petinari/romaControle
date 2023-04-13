package config

import "github.com/redis/go-redis/v9"

var ClientRedis *redis.Client

func InitRedis() {
	ClientRedis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
