package middleware

import (
	"fmt"
	"time"
	"wallet-frame/api"
	"wallet-frame/global"
	"wallet-frame/utils"

	"github.com/gin-gonic/gin"
)

func RedisRequestLockRequestLock(key string, timeout time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.GetInt64("accountId")
		if userId <= 0 {
			api.Response(ctx, 1, "redis访问锁要先用户登录", nil)
			ctx.Abort()
			return
		}

		lKey := fmt.Sprintf("lock:%s:%d", key, userId)
		lValue := fmt.Sprintf("%d", time.Now().UnixNano())
		l := utils.NewRedisLock(global.RedisC, lKey, lValue, timeout)
		err := l.TryLock()
		if err != nil {
			api.Response(ctx, 1, "访问评率太高", nil)
			ctx.Abort()
			return
		}
		ctx.Next()
		_ = l.UnLock()
	}
}
