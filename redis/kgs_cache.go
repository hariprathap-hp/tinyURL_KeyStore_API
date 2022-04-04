package keystore

import (
	dom_kgs "test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/domain"
)

type KGScache interface {
	Set(key string, value string)
	Get(key string) *dom_kgs.Key
}
