package middleware

import (
	"fmt"
	"wallet-frame/api"

	"github.com/gin-gonic/gin"
)

func RecoverMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				api.Response(c, 1, fmt.Sprint(err), nil)
				c.Abort()
				return
			}
		}()
	}
}
