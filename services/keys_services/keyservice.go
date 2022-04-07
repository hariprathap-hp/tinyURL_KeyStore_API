package keys_services

import (
	"fmt"
	dom_keys "test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/domain"
	keystore "test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/redis"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
)

var kgs_cache keystore.KGScache

func init() {
	kgs_cache = keystore.NewRedisCache("localhost:6379", 1, 10)
}

var (
	KeyService keyServicesInterface = &keyservices{}
)

type keyservices struct{}

type keyServicesInterface interface {
	Get() (*string, *errors.RestErr)
	Populate() *errors.RestErr
	Cache() *errors.RestErr
}

func (ks *keyservices) Get() (*string, *errors.RestErr) {
	fmt.Println("Inside Get services")
	for {
		if k := kgs_cache.Get(); k != "" {
			return &k, nil
		}
		//if the list is empty, cache the keys again and get key from the cache
		if err := ks.Cache(); err != nil {
			return nil, err
		}
	}
}

func (ks *keyservices) Populate() *errors.RestErr {
	var key dom_keys.Key
	if err := key.Populate("populate"); err != nil {
		return err
	}
	return nil
}

func (ks *keyservices) Cache() *errors.RestErr {
	var key dom_keys.Key
	results, err := key.Get(20, true)
	if err != nil {
		return err
	}
	var keys []string
	for _, v := range results {
		keys = append(keys, v.Token)
	}
	PopulateRedis(keys)
	return nil
}

func PopulateRedis(keys []string) {
	kgs_cache.Set(keys)
}
