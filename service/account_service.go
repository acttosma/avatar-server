package service

import (
	"acttos.com/avatar/model/api/resp"
	"acttos.com/avatar/pkg/util/logger"
)

type AccountService struct{}

// 用户登录方法
func (as *AccountService) LoginAfterWechatOAuth(openId, mobile, verifyCode string, tableId int64) (*resp.ActLogin, error) {
	logger.Monitor.Debugf("openId:%s, mobile:%s, verifyCode:%s,tableId:%d", openId, mobile, verifyCode, tableId)
	return &resp.ActLogin{}, nil
}
