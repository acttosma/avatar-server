// Package rscode
// @Description: DO NOT MODIFY THIS FILE, UNLESS YOU KNOW WHAT YOU ARE DOING
package rscode

import (
	"avatarmeta.cc/avatar/model/api/resp"
	"avatarmeta.cc/avatar/pkg/util/logger"
)

var enUS *Lang

func en_US() Lang {
	if enUS == nil {
		initEnUS()
	}

	return *enUS
}

func initEnUS() Lang {
	once.Do(func() {
		enUS = &Lang{
			// 默认httpStatusCode相关
			RSP_CODE_SUCCEED:         resp.Base{Code: 200, Message: "SUCCEED"},
			RSP_CODE_NOT_FOUND_ERROR: resp.Base{Code: 404, Message: "Page not found"},

			// 参数错误相关
			RSP_CODE_PARAM_ERROR:             resp.Base{Code: -1001, Message: "Parameters illegal"},
			RSP_CODE_SYSTEM_ERROR:            resp.Base{Code: -1002, Message: "System Error,please try again later"},
			RSP_CODE_OPERATION_FAILED_ERROR:  resp.Base{Code: -1003, Message: "Operation failed"},
			RSP_CODE_DATA_INCOMPLETE_ERROR:   resp.Base{Code: -1004, Message: "Data incomplete"},
			RSP_CODE_DATA_NOT_EXIST_ERROR:    resp.Base{Code: -1005, Message: "Date not exist"},
			RSP_CODE_OPERATION_DENIED_ERROR:  resp.Base{Code: -1006, Message: "Operation denied"},
			RSP_CODE_PERMISSION_DENIED_ERROR: resp.Base{Code: -1007, Message: "Permission denied"},

			// account模块错误相关
			RSP_CODE_ACCOUNT_NOT_LOGIN_ERROR:         resp.Base{Code: -2001, Message: "Illegal access, need login"},
			RSP_CODE_ACCOUNT_WECHAT_ERROR:            resp.Base{Code: -2002, Message: "Wechat error,check the parameters"},
			RSP_CODE_ACCOUNT_NO_OPENID_ERROR:         resp.Base{Code: -2003, Message: "Must access within WeChat App"},
			RSP_CODE_ACCOUNT_NOT_EXIST_ERROR:         resp.Base{Code: -2004, Message: "Account not exist"},
			RSP_CODE_ACCOUNT_PWD_INVALID_ERROR:       resp.Base{Code: -2005, Message: "Password incorrect"},
			RSP_CODE_ACCOUNT_PWD_ALREADY_SET_ERROR:   resp.Base{Code: -2006, Message: "Password has been set already,do not try again"},
			RSP_CODE_ACCOUNT_MEM_CARD_EMPTY_ERROR:    resp.Base{Code: -2007, Message: "Member card not found"},
			RSP_CODE_ACCOUNT_INVITER_NOT_EXIST_ERROR: resp.Base{Code: -2008, Message: "The inviter does not exist"},
			RSP_CODE_ACCOUNT_ALREADY_EXIST_ERROR:     resp.Base{Code: -2009, Message: "The account is already registered"},

			// captcha模块错误相关
			RSP_CODE_CAPTCHA_INVALID_ERROR:       resp.Base{Code: -5001, Message: "The captcha is incorrect,please re-enter"},
			RSP_CODE_CAPTCHA_NONCE_INVALID_ERROR: resp.Base{Code: -5002, Message: "Illegal access,please be careful"},

			// qrcode模块错误相关
			RSP_CODE_QRCODE_GEN_ERROR: resp.Base{Code: -7001, Message: "System error,please try again later"},
		}
		logger.Monitor.Debug("en_US has been initialized")
	})

	return *enUS
}
