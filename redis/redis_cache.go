package keystore

import (
	"encoding/json"
	"fmt"
	dom_kgs "test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/domain"
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

func (cache *redisCache) Set(key string, value string) {
	client := cache.getClient()

	fmt.Println("Setting Keys in Redis")
	fmt.Println(key, value)
	client.Set(key, value, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *dom_kgs.Key {
	client := cache.getClient()
	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}
	res := dom_kgs.Key{}
	json.Unmarshal([]byte(val), &res)
	return &res
}
