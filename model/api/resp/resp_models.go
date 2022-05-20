package resp

// 基本响应协议
type Base struct {
	// 返回码,0表示正常,其它值表示出现问题
	Code int `json:"code"`
	// 对Code的简要描述
	Message string `json:"msg"`
}

type ActLogin struct {
	// 账户ID 老用户情况下返回此值,新用户尚需绑定手机号码
	AccountId int64 `json:"accountId"`
	// 令牌 老用户情况下返回此值,后续响应可携带鉴权用
	Token string `json:"token"`
	// 第三方开放ID
	OpenId string `json:"openId"`
	// 是否已经设置密码, true:已设置,fale:未设置
	IsPwdSet bool `json:"isPwdSet"`
}

type CaptchaGet struct {
	// 图片验证码对应的唯一随机串
	Nonce string `json:"nonce"`
	// 验证码图片,含图片bytes的BASE64编码图片,格式png
	Captcha string `json:"captcha"`
}

type CheckCaptcha struct {
	// 图片验证码校验码对应的Key
	CaptchaNonceKey string `json:"captchaNonceKey"`
	// 图片验证码校验码,验证成功情况下会返回此值,单次有效,用于在需要使用图片验证码防止机刷接口调用
	CaptchaNonce string `json:"captchaNonce"`
}

type QRCodeImg struct {
	// 二维码图片,图片bytes的BASE64编码图片,格式png
	Image string `json:"image"`
}
