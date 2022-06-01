package routers

import (
	"avatarmeta.cc/avatar/service"
)

var (
	accountService = new(service.AccountService)
	captchaService = new(service.CaptchaService)
)
