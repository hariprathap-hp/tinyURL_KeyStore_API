package keys

import (
	"fmt"
	"net/http"
	"test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/services/keys_services"

	"github.com/gin-gonic/gin"
)

func GetKey(c *gin.Context) {
	s, err := keys_services.KeyService.Get()
	if err != nil {
		fmt.Println("JSON Error")
		c.JSON(err.Status, err)
		return
	}
	fmt.Println("Returning Key -- ", s)
	c.JSON(http.StatusOK, s)
}

func PopulateKeys(c *gin.Context) {
	err := keys_services.KeyService.Populate()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, "Keys are Populated")
}

func CacheKeys(c *gin.Context) {
	err := keys_services.KeyService.Cache()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, "Values cached")
}
