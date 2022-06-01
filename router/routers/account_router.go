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

	code, resp := accountService.Register(mail, password, inviteCode)
	c.JSON(code, resp)
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

	code, resp := accountService.LoginWithMail(mail, password)
	c.JSON(code, resp)
}

// @BasePath /api/v1
// @Tags Audience - Account Module
// @Summary User change the password while logged in
// @Description Change the login password, remember user must be online when do this action, otherwise please see '/account/resetPassword'
// @Accept json
// @Produce json
// @Param Authorization	header string true	"The JWT (called 'authorization' in the return value) after user logged in"
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

// @BasePath /api/v1
// @Tags Audience - Account Module
// @Summary User change the trade password while logged in
// @Description Change the trade password, remember user must be online when do this action
// @Accept json
// @Produce json
// @Param Authorization	header string true	"The JWT (called 'authorization' in the return value) after user logged in"
// @Param data body req.ChangePassword true "ChangePassword"
// @Success 200 {object} resp.Base
// @Router /account/changeTradePwd [POST]
func (ar *AccountRouter) ChangeTradePassword(c *gin.Context) {
	var reqModel req.ChangePassword
	if err := c.ShouldBind(&reqModel); err != nil {
		logger.Monitor.Errorf("ChangeTradePassword param parse error:%+v", err)
		c.JSON(http.StatusBadRequest, rscode.Code(c).RSP_CODE_PARAM_ERROR)
		return
	}

	actIdStr := c.Request.Header.Get("AccountId")
	prePassword := reqModel.PreviousPassword
	password := reqModel.Password

	code, resp := accountService.ChangeTradePassword(actIdStr, prePassword, password)
	c.JSON(code, resp)
}

// @BasePath /api/v1
// @Tags Audience - Account Module
// @Summary User set the trade password while logged in
// @Description Set the trade password, remember user must be online when do this action
// @Accept json
// @Produce json
// @Param Authorization	header string true	"The JWT (called 'authorization' in the return value) after user logged in"
// @Param data body req.SetPassword true "SetPassword"
// @Success 200 {object} resp.Base
// @Router /account/setTradePwd [POST]
func (ar *AccountRouter) SetTradePassword(c *gin.Context) {
	var reqModel req.SetPassword
	if err := c.ShouldBind(&reqModel); err != nil {
		logger.Monitor.Errorf("SetTradePassword param parse error:%+v", err)
		c.JSON(http.StatusBadRequest, rscode.Code(c).RSP_CODE_PARAM_ERROR)
		return
	}

	actIdStr := c.Request.Header.Get("AccountId")
	password := reqModel.Password

	code, resp := accountService.SetTradePassword(actIdStr, password)
	c.JSON(code, resp)
}
