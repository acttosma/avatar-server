package util

import (
	"crypto/md5"
	"crypto/rc4"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"sync"
	"time"

	"acttos.com/avatar/model/constant"
	"acttos.com/avatar/pkg/setting"
	"acttos.com/avatar/pkg/util/logger"
	"acttos.com/avatar/pkg/util/redis"
	"github.com/dgrijalva/jwt-go"
	"github.com/xxtea/xxtea-go/xxtea"
)

var CryptoHelper *Crypto
var once sync.Once

func InitCryptoHelper() {
	once.Do(func() {
		CryptoHelper = &Crypto{}
	})
}

type Crypto struct{}

func (c *Crypto) Md5Digest(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func (c *Crypto) Sha1Digest(str string) string {
	data := []byte(str)
	has := sha1.Sum(data)
	sha1str := fmt.Sprintf("%x", has)
	return sha1str
}

func (c *Crypto) Sha256Digest(str string) string {
	data := []byte(str)
	has := sha256.Sum256(data)
	sha256str := fmt.Sprintf("%x", has)
	return sha256str
}

func (c *Crypto) Base64EncodingWithBytes(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

func (c *Crypto) Base64Encoding(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func (c *Crypto) Base64Decoding(str string) string {
	sDec, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println("Failed to base64-decoding str:", str)
	}

	return string(sDec)
}

func (c *Crypto) RC4EncryptWithBase64(plain, key string) string {
	cipherBytes := c.RC4(plain, key)
	cipherString := c.Base64Encoding(string(cipherBytes))

	return cipherString
}

func (c *Crypto) RC4DecryptWithBase64(cipher, key string) string {
	cipher = c.Base64Decoding(cipher)
	plainBytes := c.RC4(cipher, key)
	plainString := string(plainBytes)

	return plainString
}

func (c *Crypto) RC4(source, key string) []byte {
	var keyBytes []byte = []byte(key)
	rc4Cipher, err := rc4.NewCipher(keyBytes)
	if err != nil {
		log.Println("Failed to generate new RC4 Cipher with key:", key)
		return []byte{}
	}
	sourceBytes := []byte(source)
	cipherBytes := make([]byte, len(sourceBytes))
	rc4Cipher.XORKeyStream(cipherBytes, sourceBytes)

	return cipherBytes
}

type CustomClaims struct {
	AccountId   string `json:"a"`
	AccountType string `json:"b"`
	jwt.StandardClaims
}

// 根据用户的用户名和密码产生token
func (c *Crypto) GenerateJWT(accountId, accountType string) (string, error) {
	//设置token有效时间,720小时候后过期(30天*24小时/天)
	nowTime := time.Now()
	expireTime := nowTime.Add(720 * time.Hour)

	issuer := xxtea.EncryptString(accountId+"@"+setting.GetInstance().Organ, setting.GetInstance().XxteaKey)
	claims := CustomClaims{
		AccountId:   accountId,
		AccountType: accountType,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: expireTime.Unix(),
		},
	}

	logger.Monitor.Debugf("JWT claims:%+v", claims)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串,再用于获取完整、已签名的token
	jwtSecret := []byte(setting.GetInstance().JwtSecret)
	token, err := tokenClaims.SignedString(jwtSecret)

	// 如想要实现随时踢人的效果,则需要redis的支持
	if accountType == constant.DEFAULT_JWT_USER_ROLE_AUDIENCE {
		redis.Helper.HSet(constant.REDIS_KEY_TOKEN_USER_MAP_KEY, accountId, token)
	} else {
		redis.Helper.HSet(constant.REDIS_KEY_TOKEN_MCH_MAP_KEY, accountId, token)
	}

	return token, err
}

// 根据传入的token值获取到Claims对象信息,(进而获取其中的用户信息)
func (c *Crypto) CheckJWT(token string) (accountId, accountType string, isLegal bool) {
	claims, err := parseJWT(token)
	if err != nil {
		logger.Monitor.Errorf("Error when parsing JWT:%s error:%+v", token, err)
		return "", "-1", false
	}
	accountId = claims.AccountId
	accountType = claims.AccountType
	logger.Monitor.Debugf("JWT claims:%+v", claims)

	// 如想要实现随时踢人的效果,则需要redis的支持
	var tokenInRedis string
	if accountType == constant.DEFAULT_JWT_USER_ROLE_AUDIENCE {
		tokenInRedis = redis.Helper.HGet(constant.REDIS_KEY_TOKEN_USER_MAP_KEY, accountId)
	} else {
		tokenInRedis = redis.Helper.HGet(constant.REDIS_KEY_TOKEN_MCH_MAP_KEY, accountId)
	}
	if len(tokenInRedis) == 0 || token != tokenInRedis {
		return "", "-1", false
	}

	issuer, err := xxtea.DecryptString(claims.Issuer, setting.GetInstance().XxteaKey)
	if err != nil || issuer != accountId+"@"+setting.GetInstance().Organ {
		return "", "-1", false
	}

	return accountId, accountType, true
}

func (c *Crypto) GenerateSaltedPassword(password, pwdSalt string) string {
	return c.Md5Digest(password + "_" + pwdSalt)
}

// 根据传入的token值获取到Claims对象信息,(进而获取其中的用户名和密码)
func parseJWT(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := []byte(setting.GetInstance().JwtSecret)
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象,并使用断言,将该对象转换为我们自己定义的Claims
		// 要传入指针,项目中结构体用指针传递,可节省空间。
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
