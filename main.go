package main

import (
	"strconv"

	"avatarmeta.cc/avatar/pkg/midlwre"
	"avatarmeta.cc/avatar/pkg/setting"
	"avatarmeta.cc/avatar/pkg/util/logger"
	"avatarmeta.cc/avatar/pkg/util/startup"
	"avatarmeta.cc/avatar/router"
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
