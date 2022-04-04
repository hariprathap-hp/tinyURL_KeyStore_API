package keys_services

import (
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
	dom_keys "test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/domain"
	keystore "test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/redis"
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
	var key dom_keys.Key
	result, err := key.Get(1, false)
	if err != nil {
		return nil, err
	}
	return &result[0].Token, nil
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

func PopulateRedis(keys []string) *errors.RestErr {
	for _, v := range keys {
		kgs_cache.Set(v, v+"a")
	}
	return nil
}
