package service

import (
	"acttos.com/avatar/model/api/resp"
	"acttos.com/avatar/pkg/util/logger"
)

type AccountService struct{}

// 用户登录方法
func (as *AccountService) Register(mail, password, inviteCode string) (*resp.ActRegister, error) {
	logger.Monitor.Debugf("mail:%s, password:%s, inviteCode:%s", mail, password, inviteCode)
	return &resp.ActRegister{}, nil
}

// 用户登录方法
func (as *AccountService) LoginWithMail(mail, password string) (*resp.ActLogin, error) {
	logger.Monitor.Debugf("mail:%s, password:%s", mail, password)
	return &resp.ActLogin{}, nil
}
