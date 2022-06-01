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
	savedAct, err := account.Add()
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
