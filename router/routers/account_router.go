package routers

import (
	"acttos.com/avatar/model/api/resp/rscode"
	"net/http"

	"acttos.com/avatar/model/api/req"
	"acttos.com/avatar/pkg/util/logger"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

// @BasePath /api/v1
// @Tags 普通用户-账号相关
// @Summary 用户登录使用
// @Description 此接口用于在用户登录使用
// @Accept json
// @Produce json
// @Param data body req.ActLogin true "ActLogin"
// @Success 200 {object} resp.ActLogin
// @Router /account/login [POST]
func (ar *AccountRouter) LoginAfterWechatOAuth(c *gin.Context) {
	var reqModel req.ActLogin
	if err := c.ShouldBind(&reqModel); err != nil {
		logger.Monitor.Errorf("LoginWithMobile param parse error:%+v", err)
		c.JSON(http.StatusBadRequest, rscode.Code(c).RSP_CODE_PARAM_ERROR)
		return
	}

	openId := reqModel.OpenId
	mobile := reqModel.Mobile
	verifyCode := reqModel.VerifyCode

	loginResp, err := accountService.LoginAfterWechatOAuth(openId, mobile, verifyCode, reqModel.TableId)
	if err != nil {
		logger.Monitor.Errorf("Error:%+v", err)
		c.JSON(http.StatusInternalServerError, rscode.Code(c).RSP_CODE_SYSTEM_ERROR)
		return
	}

	c.JSON(http.StatusOK, loginResp)
}
