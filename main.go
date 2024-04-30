package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"wallet-frame/global"
	_ "wallet-frame/logger"
	"wallet-frame/routers"
)

var envString string

func main() {
	flag.StringVar(&envString, "env", "dev", "请输入当前环境配置文件")
	flag.Parse()
	fmt.Println(envString, "当前环境")

	// 1.初始化全局组件
	global.Initialize(envString)

	//2.初始化路由
	router := routers.InitRouters()

	// 获取端口号
	PORT := strconv.Itoa(global.Config.Port)
	fmt.Println(PORT + "当前端口")

	go func() {
		// 启动服务
		if err := router.Run(fmt.Sprintf(":%s", PORT)); err != nil {
			fmt.Println(fmt.Sprintf("服务启动失败:%s", err.Error()))
		}
		fmt.Println(fmt.Sprintf("服务已经启动:localhost:%s", PORT))
	}()
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
