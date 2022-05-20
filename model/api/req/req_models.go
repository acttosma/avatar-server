package req

// 用户使用手机号码+验证码登录请求参数
type ActLogin struct {
	// 桌台Id [必填]
	TableId int64 `form:"tableId" json:"tableId" binding:"required"`
	// 开放平台ID,支持微信,支付宝,微博等平台用户 [必填]
	OpenId string `form:"openId" json:"openId" uri:"openId" xml:"openId" binding:"required"`
	// 手机号码 [必填]
	Mobile string `form:"mobile" json:"mobile" uri:"mobile" xml:"mobile" binding:"required"`
	// 手机验证码 [必填]
	VerifyCode string `form:"verifyCode" json:"verifyCode" uri:"verifyCode" xml:"verifyCode" binding:"required"`
}
