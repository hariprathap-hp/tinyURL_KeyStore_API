package keystore

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) KGScache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	fmt.Println("Host is -- ", cache.host)
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(keys []string) {
	client := cache.getClient()
	fmt.Println("Setting Keys in a Redis List")
	client.LPush("kgs", keys)
}

func (cache *redisCache) Get() []string {
	client := cache.getClient()
	result := client.LRange("kgs", 0, -1)
	client.Del("kgs")
	return result.Val()
}
