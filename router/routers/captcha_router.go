package routers

import (
	"acttos.com/avatar/model/api/resp"
	"acttos.com/avatar/model/api/resp/rscode"
	"acttos.com/avatar/pkg/util/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CaptchaRouter struct{}

// @BasePath /api/v1
// @Tags 不限用户-图片验证码相关
// @Summary 图片验证码生成接口
// @Description 根据给定的参数生成图片验证码,为提高安全性,不建议图片宽高太大或字符个数太少
// @Accept text/plain
// @Produce json
// @Param w query int false "图片验证码的宽,默认值100,最大值500,单位:px"
// @Param h query int false "图片验证码的高,默认值30,最大值200,单位:px"
// @Param l query int false "图片验证码包含字符的个数,范围[4,8],请根据宽高显示的实际效果来决定此参数,默认值4,单位:个"
// @Success 200 {object} resp.CaptchaGet
// @Router /captcha/get [GET]
func (cr *CaptchaRouter) Get(c *gin.Context) {
	w := c.Query("w")
	h := c.Query("h")
	l := c.Query("l")

	nonce, captcha, err := captchaService.GenerateCaptcha(w, h, l)
	if err != nil {
		logger.Monitor.Errorf("Error:%+v", err)
		c.JSON(http.StatusInternalServerError, rscode.Code(c).RSP_CODE_SYSTEM_ERROR)
		return
	}

	c.JSON(http.StatusOK, resp.CaptchaGet{
		Nonce:   nonce,
		Captcha: captcha,
	})
}

// @BasePath /api/v1
// @Tags 不限用户-图片验证码相关
// @Summary 图片验证码校验接口
// @Description 接收图片验证码参数对输入的captcha对儿进行校验
// @Accept text/plain
// @Produce json
// @Param ci query string true "图片验证码的ID(captchaId)"
// @Param cc query string true "图片验证码显示的字符(captchaCode),大小写不敏感"
// @Success 200 {object} resp.CheckCaptcha
// @Router /captcha/check [GET]
func (cr *CaptchaRouter) Check(c *gin.Context) {
	captchaId := c.Query("ci")
	captchaCode := c.Query("cc")

	check, err := captchaService.CheckCaptcha(captchaId, captchaCode)
	if err != nil {
		logger.Monitor.Errorf("Error:%+v", err)
		c.JSON(http.StatusInternalServerError, rscode.Code(c).RSP_CODE_SYSTEM_ERROR)
		return
	}
	c.JSON(http.StatusOK, check)
}
