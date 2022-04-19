package keys_services

import (
	dom_keys "test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/domain"
	"test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/kgs_redis"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
	zlogger "test3/hariprathap-hp/system_design/utils_repo/log_utils"
)

var (
	KeyService keyServicesInterface = &keyservices{}
)

type keyservices struct{}

type keyServicesInterface interface {
	Get() ([]string, *errors.RestErr)
	Populate() *errors.RestErr
	Cache() *errors.RestErr
}

func (ks *keyservices) Get() ([]string, *errors.RestErr) {
	zlogger.Info("service keystore: func Get(): getting keys to be distributed to the clients")
	var key dom_keys.Key
	results, err := key.Get(25, true)
	if err != nil {
		return nil, err
	}
	return results, nil
	/*for {
		if k := kgs_redis.Rcache.Get(); len(k) != 0 {
			//renew the local cache once again after handing over the available keys to app_cache
			if err := ks.Cache(); err != nil {
				return nil, err
			}
			return k, nil
		}
		//if the list is empty, cache the keys again and get key from the cache
		if err := ks.Cache(); err != nil {
			return nil, err
		}
	}*/
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
	results, err := key.Get(25, true)
	if err != nil {
		return err
	}

	PopulateRedis(results)
	return nil
}

func PopulateRedis(keys []string) {
	kgs_redis.Rcache.Set(keys)
}
