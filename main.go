package main

import (
	"strconv"

	"acttos.com/avatar/pkg/midlwre"
	"acttos.com/avatar/pkg/setting"
	"acttos.com/avatar/pkg/util/logger"
	"acttos.com/avatar/pkg/util/startup"
	"acttos.com/avatar/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 必须优先调用初始化方法
	startup.Initialize()

	gin.SetMode(setting.GetInstance().ServerMode)

	// r := gin.Default()
	r := customDefault()
	// 输入输出参数日志
	r.Use(midlwre.ParamLog())

	router.Init(r)

	r.Run(":" + strconv.Itoa(setting.GetInstance().HttpPort))
}

func customDefault() *gin.Engine {
	engine := gin.New()
	engine.Use(logger.GinLogger(), logger.GinRecovery(true))
	return engine
}
