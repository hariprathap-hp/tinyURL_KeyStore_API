package kgs_redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var (
	Rcache kgscacheinterface = &redisCache{}
	client *redis.Client
)

func init() {
	kgs_cache := NewRedisCache("localhost:6379", 1, 10)
	client = redis.NewClient(&redis.Options{
		Addr:     kgs_cache.host,
		Password: "",
		DB:       kgs_cache.db,
	})
}

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

type kgscacheinterface interface {
	Set([]string)
	Get() []string
}

func NewRedisCache(host string, db int, exp time.Duration) redisCache {
	return redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) Set(keys []string) {
	fmt.Println("Setting Keys in a Redis List")
	client.LPush("kgs", keys)
}

func (cache *redisCache) Get() []string {
	result := client.LRange("kgs", 0, -1)
	client.Del("kgs")
	return result.Val()
}
