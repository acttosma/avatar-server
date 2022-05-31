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
	docs.SwaggerInfo.Title = "AvatarMeta interaction between the server side and web side"
	docs.SwaggerInfo.Description = "All the interfaces on this page have NOT tested deeply,if you have any question,please <a href='mailto:acttosma@126.com'>let me know</a>.<br/>All the interfaces should access with BASE URL as the prefix: /api/v1,<br/>For example: the whole URL of login interface: https://host:port/api/v1/account/login"

	gin.GET("/api/v1/swagger/*any", midlwre.SwaggerCheck(), ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := gin.Group("/api/v1")

	account := v1.Group("/account")
	{
		account.POST("/register", accountRouter.Register)
		account.POST("/login", accountRouter.LoginWithMail)
	}

	captcha := v1.Group("/captcha")
	{
		captcha.GET("/get", captchaRouter.Get)
		captcha.GET("/check", captchaRouter.Check)
	}

	qrcode := v1.Group("/qrcode")
	{
		qrcode.GET("/gen", qrcodeRouter.Gen)
	}
}
