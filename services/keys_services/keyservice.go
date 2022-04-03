package keys_services

import (
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
	dom_keys "test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/domain"
)

var (
	KeyService keyServicesInterface = &keyservices{}
)

type keyservices struct{}

type keyServicesInterface interface {
	Get() (*string, *errors.RestErr)
	Populate() *errors.RestErr
}

func (ks *keyservices) Get() (*string, *errors.RestErr) {
	var key dom_keys.Key
	result, err := key.Get(1)
	if err != nil {
		return nil, err
	}
	return &result.Token, nil
}

func (ks *keyservices) Populate() *errors.RestErr {
	var key dom_keys.Key
	if err := key.Populate("populate"); err != nil {
		return err
	}
	return nil
}
