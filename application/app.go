package application

import (
	"test3/hariprathap-hp/system_design/tinyURL/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("About to Start the Application")
	router.Run(":8085")
}
