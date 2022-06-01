package routers

import (
	"avatarmeta.cc/avatar/model/api/resp/rscode"
	"net/http"

	"avatarmeta.cc/avatar/model/api/req"
	"avatarmeta.cc/avatar/pkg/util/logger"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

// @BasePath /api/v1
// @Tags Audience - Account Module
// @Summary User register as an account through this api
// @Description Provide mail,password and invite code to register as a new account
// @Accept json
// @Produce json
// @Param data body req.ActRegister true "ActRegister"
// @Success 200 {object} resp.ActRegister
// @Router /account/register [POST]
func (ar *AccountRouter) Register(c *gin.Context) {
	var reqModel req.ActRegister
	if err := c.ShouldBind(&reqModel); err != nil {
		logger.Monitor.Errorf("Register param parse error:%+v", err)
		c.JSON(http.StatusBadRequest, rscode.Code(c).RSP_CODE_PARAM_ERROR)
		return
	}

	mail := reqModel.Mail
	password := reqModel.Password
	inviteCode := reqModel.InviteCode

	resp, err := accountService.Register(mail, password, inviteCode)
	if err != nil {
		logger.Monitor.Errorf("Error:%+v", err)
		c.JSON(http.StatusInternalServerError, rscode.Code(c).RSP_CODE_SYSTEM_ERROR)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @BasePath /api/v1
// @Tags Audience - Account Module
// @Summary User login with mail and password
// @Description User login action  with resp.ActLogin returned
// @Accept json
// @Produce json
// @Param data body req.ActLogin true "ActLogin"
// @Success 200 {object} resp.ActLogin
// @Router /account/login [POST]
func (ar *AccountRouter) LoginWithMail(c *gin.Context) {
	var reqModel req.ActLogin
	if err := c.ShouldBind(&reqModel); err != nil {
		logger.Monitor.Errorf("LoginWithMobile param parse error:%+v", err)
		c.JSON(http.StatusBadRequest, rscode.Code(c).RSP_CODE_PARAM_ERROR)
		return
	}

	mail := reqModel.Mail
	password := reqModel.Password

	loginResp, err := accountService.LoginWithMail(mail, password)
	if err != nil {
		logger.Monitor.Errorf("Error:%+v", err)
		c.JSON(http.StatusInternalServerError, rscode.Code(c).RSP_CODE_SYSTEM_ERROR)
		return
	}

	c.JSON(http.StatusOK, loginResp)
}

// @BasePath /api/v1
// @Tags Audience - Account Module
// @Summary User change the password while logged in
// @Description Change the login password, remember user must be online when do this action, otherwise please see '/account/resetPassword'
// @Accept json
// @Produce json
// @Param data body req.ChangePassword true "ChangePassword"
// @Success 200 {object} resp.Base
// @Router /account/changePwd [POST]
func (ar *AccountRouter) ChangePassword(c *gin.Context) {
	var reqModel req.ChangePassword
	if err := c.ShouldBind(&reqModel); err != nil {
		logger.Monitor.Errorf("ChangePassword param parse error:%+v", err)
		c.JSON(http.StatusBadRequest, rscode.Code(c).RSP_CODE_PARAM_ERROR)
		return
	}

	actIdStr := c.Request.Header.Get("AccountId")
	prePassword := reqModel.PreviousPassword
	password := reqModel.Password

	code, resp := accountService.ChangePassword(actIdStr, prePassword, password)
	c.JSON(code, resp)
}
