package application

import "test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/controller/keys"

func mapUrls() {
	router.GET("/cachekeys", keys.CacheKeys)
	router.GET("/getkey", keys.GetKey)
	router.POST("/populate", keys.PopulateKeys)
}
