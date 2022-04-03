package keys

import (
	"net/http"
	"test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/services/keys_services"

	"github.com/gin-gonic/gin"
)

func GetKey(c *gin.Context) {
	s, err := keys_services.KeyService.Get()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, *s)
}

func PopulateKeys(c *gin.Context) {
	err := keys_services.KeyService.Populate()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, "Keys are Populated")
}
