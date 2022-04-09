package application

import (
	zlogger "test3/hariprathap-hp/system_design/utils_repo/log_utils"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	zlogger.Info("About to start the service")
	router.Run(":8081")
}
