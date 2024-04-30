package api

import (
	"wallet-frame/logger"

	"github.com/gin-gonic/gin"
)

func TestApi(c *gin.Context) {
	logger.Logs.Debug("Debuging...")
	logger.Logs.Info("Test api request...")
	logger.Logs.Warn("Warning...")
	logger.Logs.Error("Error...")
	Response(c, SUCCESS, "", nil)
}
