package service

import (
	"avatarmeta.cc/avatar/model/api/resp"
	"avatarmeta.cc/avatar/model/api/resp/rscode"
	"avatarmeta.cc/avatar/model/constant"
	"avatarmeta.cc/avatar/model/entity"
	"avatarmeta.cc/avatar/pkg/util"
	"avatarmeta.cc/avatar/pkg/util/logger"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type AccountService struct{}

func (as *AccountService) Register(mail, password, inviteCode string) (int, interface{}) {
	logger.Monitor.Debugf("mail:%s, password:%s, inviteCode:%s", mail, password, inviteCode)
	var account entity.Account

	// Check the inviter
	inviter, err := account.FindByInviteCode(inviteCode)
	if err != nil {
		logger.Monitor.Errorf("Error:%+v", err)
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}
	if inviter == nil {
		logger.Monitor.Errorf("The inviter does NOT exist")
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_ACCOUNT_NOT_EXIST_ERROR
	}

	// Check the mail
	mailAccount, err := account.FindByMail(mail)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Monitor.Errorf("Error:%+v", err)
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}
	if mailAccount != nil {
		logger.Monitor.Errorf("The mail does exist, can NOT register again")
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_ACCOUNT_ALREADY_EXIST_ERROR
	}

	// Save new account
	passwordSalt := util.TextHelper.RandomASCII(16)
	password = util.CryptoHelper.GenerateSaltedPassword(password, passwordSalt)
	account = entity.Account{
		Mail:         mail,
		Password:     password,
		PasswordSalt: passwordSalt,
		InviterId:    inviter.Id,
	}
	savedAct, err := account.Register()
	if err != nil {
		logger.Monitor.Errorf("Error:%+v", err)
		return http.StatusInternalServerError, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}

	// Handle the return response
	jwtToken, err := util.CryptoHelper.GenerateJWT(strconv.FormatInt(savedAct.Id, 10), constant.DEFAULT_JWT_USER_ROLE_AUDIENCE)
	if err != nil {
		logger.Monitor.Errorf("Error occurs when generating TOKEN:%+v", err)
		return http.StatusInternalServerError, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}

	return http.StatusOK, resp.ActRegister{
		AccountId:     savedAct.Id,
		Authorization: jwtToken,
	}
}

func (as *AccountService) LoginWithMail(mail, password string) (int, interface{}) {
	logger.Monitor.Debugf("mail:%s, password:%s", mail, password)
	var account entity.Account

	// Check the mail
	mailAccount, err := account.FindByMail(mail)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Monitor.Errorf("Error:%+v", err)
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}
	if mailAccount == nil {
		logger.Monitor.Errorf("The account does NOT exist")
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_ACCOUNT_NOT_EXIST_ERROR
	}

	// Handle the return response
	password = util.CryptoHelper.GenerateSaltedPassword(password, mailAccount.PasswordSalt)
	if password != mailAccount.Password {
		logger.Monitor.Errorf("The password is incorrect")
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_ACCOUNT_PWD_INVALID_ERROR
	}

	jwtToken, err := util.CryptoHelper.GenerateJWT(strconv.FormatInt(mailAccount.Id, 10), constant.DEFAULT_JWT_USER_ROLE_AUDIENCE)
	if err != nil {
		logger.Monitor.Errorf("Error occurs when generating TOKEN:%+v", err)
		return http.StatusInternalServerError, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}

	return http.StatusOK, resp.ActLogin{
		AccountId:     mailAccount.Id,
		Authorization: jwtToken,
	}
}

func (as *AccountService) ChangePassword(actIdStr, prePassword, password string) (int, interface{}) {
	actId, err := strconv.ParseInt(actIdStr, 10, 64)
	if err != nil {
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_PARAM_ERROR
	}

	var account entity.Account
	// Check the mail
	act, err := account.FindById(actId)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Monitor.Errorf("Error:%+v", err)
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}
	if act == nil {
		logger.Monitor.Errorf("The account does NOT exist")
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_ACCOUNT_NOT_EXIST_ERROR
	}

	// Check the previous password
	prePassword = util.CryptoHelper.GenerateSaltedPassword(prePassword, act.PasswordSalt)
	if prePassword != act.Password {
		logger.Monitor.Errorf("The password is incorrect")
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_ACCOUNT_PWD_INVALID_ERROR
	}

	// Handle the return response
	act.PasswordSalt = util.TextHelper.RandomASCII(16)
	password = util.CryptoHelper.GenerateSaltedPassword(password, act.PasswordSalt)
	act.Password = password
	act, err = act.Save()
	if err != nil {
		logger.Monitor.Errorf("Error:%+v", err)
		return http.StatusInternalServerError, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}

	return http.StatusOK, rscode.Eng().RSP_CODE_SUCCEED
}

func (as *AccountService) ChangeTradePassword(actIdStr, prePassword, password string) (int, interface{}) {
	actId, err := strconv.ParseInt(actIdStr, 10, 64)
	if err != nil {
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_PARAM_ERROR
	}

	var account entity.Account
	// Check the mail
	act, err := account.FindById(actId)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Monitor.Errorf("Error:%+v", err)
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}
	if act == nil {
		logger.Monitor.Errorf("The account does NOT exist")
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_ACCOUNT_NOT_EXIST_ERROR
	}

	// Check the previous password
	prePassword = util.CryptoHelper.GenerateSaltedPassword(prePassword, act.TradePasswordSalt)
	if prePassword != act.TradePassword {
		logger.Monitor.Errorf("The trade password is incorrect")
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_ACCOUNT_PWD_INVALID_ERROR
	}

	// Handle the return response
	act.TradePasswordSalt = util.TextHelper.RandomASCII(16)
	password = util.CryptoHelper.GenerateSaltedPassword(password, act.TradePasswordSalt)
	act.TradePassword = password
	act, err = act.Save()
	if err != nil {
		logger.Monitor.Errorf("Error:%+v", err)
		return http.StatusInternalServerError, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}

	return http.StatusOK, rscode.Eng().RSP_CODE_SUCCEED
}

func (as *AccountService) SetTradePassword(actIdStr, password string) (int, interface{}) {
	actId, err := strconv.ParseInt(actIdStr, 10, 64)
	if err != nil {
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_PARAM_ERROR
	}

	var account entity.Account
	// Check the mail
	act, err := account.FindById(actId)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Monitor.Errorf("Error:%+v", err)
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}
	if act == nil {
		logger.Monitor.Errorf("The account does NOT exist")
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_ACCOUNT_NOT_EXIST_ERROR
	}

	// Check the previous password
	if act.TradePassword != "" {
		logger.Monitor.Errorf("The trade password has already set")
		return http.StatusBadRequest, rscode.Eng().RSP_CODE_ACCOUNT_TRADE_PWD_ALREADY_SET_ERROR
	}

	// Handle the return response
	act.TradePasswordSalt = util.TextHelper.RandomASCII(16)
	password = util.CryptoHelper.GenerateSaltedPassword(password, act.TradePasswordSalt)
	act.TradePassword = password
	act, err = act.Save()
	if err != nil {
		logger.Monitor.Errorf("Error:%+v", err)
		return http.StatusInternalServerError, rscode.Eng().RSP_CODE_SYSTEM_ERROR
	}

	return http.StatusOK, rscode.Eng().RSP_CODE_SUCCEED
}
