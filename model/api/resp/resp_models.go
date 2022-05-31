package resp

// Base response protocol
type Base struct {
	// the return code, 0 means OK, other codes mean failed
	Code int `json:"code"`
	// The simple description of the code, the request should NOT use this value directly, it must be translated to another suitable message
	Message string `json:"msg"`
}

type ActRegister struct {
	// The ID of the account
	AccountId int64 `json:"accountId"`
}

type ActLogin struct {
	// The ID of the account
	AccountId int64 `json:"accountId"`
}

type CaptchaGet struct {
	// The nonce key of the captcha
	Nonce string `json:"nonce"`
	// The captcha image of request, in BASE64 format
	Captcha string `json:"captcha"`
}

type CheckCaptcha struct {
	// The nonce key of the captcha check response
	CaptchaNonceKey string `json:"captchaNonceKey"`
	// The nonce value of the captcha check response
	CaptchaNonce string `json:"captchaNonce"`
}

type QRCodeImg struct {
	// The QRCode data, in BASE64 format
	Image string `json:"image"`
}
