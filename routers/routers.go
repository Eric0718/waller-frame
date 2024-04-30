package routers

import (
	"wallet-frame/api"
	"wallet-frame/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	Router := gin.Default()
	// 注册中间件
	Router.Use(
		middleware.CorsMiddleWare(),    // 跨域的
		middleware.LoggerMiddleWare(),  // 日志
		middleware.RecoverMiddleWare(), // 异常的
	)
	// 配置全局路径
	ApiGroup := Router.Group("/api/v1")
	// 注册路由
	testRouters(ApiGroup) // 账号中心
	return Router
}

func testRouters(r *gin.RouterGroup) {
	testRouter := r.Group("testApi")
	testRouter.GET("test", api.TestApi)
}
