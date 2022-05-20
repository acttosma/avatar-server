package router

import (
	"acttos.com/avatar/pkg/midlwre"
	"github.com/gin-gonic/gin"

	"acttos.com/avatar/docs"
	"acttos.com/avatar/router/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct{}

func Init(gin *gin.Engine) {
	initV1(gin)
}

func initV1(gin *gin.Engine) {
	var (
		accountRouter = new(routers.AccountRouter)
		captchaRouter = new(routers.CaptchaRouter)
		qrcodeRouter  = new(routers.QrcodeRouter)
	)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "AvatarMeta交互API说明"
	docs.SwaggerInfo.Description = "本文所列接口部分尚未联调,如有问题,及时反馈.<br/>所有接口必须加上BASE URL: /api/v1,<br/>如login接口全地址: https://host:port/api/v1/account/login, 以此类推"

	gin.GET("/api/v1/swagger/*any", midlwre.SwaggerCheck(), ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := gin.Group("/api/v1")

	account := v1.Group("/account")
	{
		// 微信授权后，使用手机号+验证码登录
		account.POST("/login", accountRouter.LoginAfterWechatOAuth)
	}

	captcha := v1.Group("/captcha")
	{
		// 获取平台图片验证码功能
		captcha.GET("/get", captchaRouter.Get)
		// 图片验证码校验功能
		captcha.GET("/check", captchaRouter.Check)
	}

	qrcode := v1.Group("/qrcode")
	{
		// 生成二维码功能 midlwre
		qrcode.GET("/gen", qrcodeRouter.Gen)
	}
}
