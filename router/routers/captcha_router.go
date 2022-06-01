package routers

import (
	"avatarmeta.cc/avatar/model/api/resp"
	"avatarmeta.cc/avatar/model/api/resp/rscode"
	"avatarmeta.cc/avatar/pkg/util/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CaptchaRouter struct{}

// @BasePath /api/v1
// @Tags Utilities-Captcha Module
// @Summary Get a captcha image
// @Description This interface returns an image data with the given parameters, the size of the request should be considered, can NOT be too high or too low.
// @Accept text/plain
// @Produce json
// @Param w query int false "the width of the image, default value is 100, max value should no more than 500. unit:px"
// @Param h query int false "the heigh of the image, default value is 30, max value should no more than 100. unit:px"
// @Param l query int false "the length of the letters shown on the image, default value is 4, it should be in the zone of [4,8]"
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
// @Tags Utilities-Captcha Module
// @Summary Check the captchaCode
// @Description This interface checks the captcha code (called 'cc')
// @Accept text/plain
// @Produce json
// @Param ci query string true "the id of the catpcha"
// @Param cc query string true "the letters on the captcha image, case insensitive"
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
