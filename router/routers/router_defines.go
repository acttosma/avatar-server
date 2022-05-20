package routers

import (
	"acttos.com/avatar/service"
)

var (
	accountService = new(service.AccountService)
	captchaService = new(service.CaptchaService)
)
