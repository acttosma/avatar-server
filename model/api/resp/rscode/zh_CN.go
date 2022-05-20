// Package rscode
// @Description: DO NOT MODIFY THIS FILE, UNLESS YOU KNOW WHAT YOU ARE DOING
package rscode

import (
	"acttos.com/avatar/model/api/resp"
	"acttos.com/avatar/pkg/util/logger"
)

var zhCN *Lang

func zh_CN() Lang {
	if zhCN == nil {
		initZhCN()
	}

	return *zhCN
}

func initZhCN() {
	once.Do(func() {
		zhCN = &Lang{
			// 默认httpStatusCode相关
			RSP_CODE_SUCCEED:         resp.Base{Code: 200, Message: "成功"},
			RSP_CODE_NOT_FOUND_ERROR: resp.Base{Code: 404, Message: "页面未找到"},

			// 参数错误相关
			RSP_CODE_PARAM_ERROR:             resp.Base{Code: -1001, Message: "请求参数有误"},
			RSP_CODE_SYSTEM_ERROR:            resp.Base{Code: -1002, Message: "系统异常,请稍后再试"},
			RSP_CODE_OPERATION_FAILED_ERROR:  resp.Base{Code: -1003, Message: "处理失败,请稍后再试"},
			RSP_CODE_DATA_INCOMPLETE_ERROR:   resp.Base{Code: -1004, Message: "您请求的数据有误,处理失败"},
			RSP_CODE_DATA_NOT_EXIST_ERROR:    resp.Base{Code: -1005, Message: "没有检索到相关数据"},
			RSP_CODE_OPERATION_DENIED_ERROR:  resp.Base{Code: -1006, Message: "对不起,您没有权限执行此操作"},
			RSP_CODE_PERMISSION_DENIED_ERROR: resp.Base{Code: -1007, Message: "对不起,您没有权限访问此页面"},

			// account模块错误相关
			RSP_CODE_ACCOUNT_NOT_LOGIN_ERROR:       resp.Base{Code: -2001, Message: "请求非法,请登录后访问"},
			RSP_CODE_ACCOUNT_WECHAT_ERROR:          resp.Base{Code: -2002, Message: "微信返回异常,请检查参数"},
			RSP_CODE_ACCOUNT_NO_OPENID_ERROR:       resp.Base{Code: -2003, Message: "请求非法,请在微信内访问"},
			RSP_CODE_ACCOUNT_NOT_EXIST_ERROR:       resp.Base{Code: -2004, Message: "用户不存在"},
			RSP_CODE_ACCOUNT_PWD_INVALID_ERROR:     resp.Base{Code: -2005, Message: "密码不正确"},
			RSP_CODE_ACCOUNT_PWD_ALREADY_SET_ERROR: resp.Base{Code: -2006, Message: "请求被拒绝,密码前期已设置"},
			RSP_CODE_ACCOUNT_MEM_CARD_EMPTY_ERROR:  resp.Base{Code: -2007, Message: "没有检索到会员卡"},

			// captcha模块错误相关
			RSP_CODE_CAPTCHA_INVALID_ERROR:       resp.Base{Code: -5001, Message: "图形验证码错误或已失效,请点击更换"},
			RSP_CODE_CAPTCHA_NONCE_INVALID_ERROR: resp.Base{Code: -5002, Message: "非法请求,请遵守法律法规"},

			// qrcode模块错误相关
			RSP_CODE_QRCODE_GEN_ERROR: resp.Base{Code: -7001, Message: "生成图片失败,请稍后再试"},
		}
		logger.Monitor.Debug("zh_CN has been initialized")
	})
}
