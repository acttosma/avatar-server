package constant

// WARNING DO NOT MODIFY, ONLY ADD
const (
	COMPACT_DAY_FORMAT_LAYOUT  = "20060102"
	DEFAULT_DAY_FORMAT_LAYOUT  = "2006-01-02"
	DEFAULT_TIME_FORMAT_LAYOUT = "2006-01-02 15:04:05"

	DEFAULT_JWT_USER_ROLE_AUDIENCE = "ACT"

	// the captcha check source value in redis
	REDIS_KEY_USER_CAPTCHA_KEY       = "captcha:captcha_key:%s"
	REDIS_KEY_USER_CAPTCHA_NONCE_KEY = "captcha:nonce_key:%s"

	REDIS_KEY_TOKEN_USER_MAP_KEY = "token:user:map"
	REDIS_KEY_TOKEN_MCH_MAP_KEY  = "token:mch:map"
)
