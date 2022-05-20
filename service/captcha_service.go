package service

import (
	"acttos.com/avatar/model/constant"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"acttos.com/avatar/model/api/resp"
	"acttos.com/avatar/pkg/util"
	"acttos.com/avatar/pkg/util/redis"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type CaptchaService struct{}

func (cs *CaptchaService) GenerateCaptcha(w, h, l string) (nonce, captcha string, err error) {
	width, err := strconv.Atoi(w)
	if err != nil || width > 500 {
		width = 100
	}

	height, err := strconv.Atoi(h)
	if err != nil || height > 200 {
		height = 30
	}

	length, err := strconv.Atoi(l)
	if err != nil || length < 4 || length > 8 {
		length = 4
	}

	if len(nonce) == 0 {
		randKey := time.Now().UnixMilli()
		nonce = util.CryptoHelper.Md5Digest(strconv.FormatInt(randKey, 10))
	}
	captchaKey := nonce
	randCode := getRandStr(length)
	redisSaveCaptcha(captchaKey, randCode, 1*time.Minute)

	bytes := imgText(width, height, randCode)

	return nonce, "data:image/png;base64," + base64.StdEncoding.EncodeToString(bytes), nil
}

func (cs *CaptchaService) CheckCaptcha(captchaId, captchaCode string) (*resp.CheckCaptcha, error) {
	if redisCheckCaptcha(captchaId, captchaCode) {
		nonce := util.CryptoHelper.Md5Digest(util.TextHelper.RandomASCII(6))
		nonceKey := util.CryptoHelper.Md5Digest(util.TextHelper.RandomASCII(6))
		if cs.SaveCaptchaNonce(nonceKey, nonce, 1*time.Minute) {
			return &resp.CheckCaptcha{
				CaptchaNonceKey: nonceKey,
				CaptchaNonce:    nonce,
			}, nil
		} else {
			return nil, errors.New("system error")
		}
	}

	return nil, errors.New("unknown error")
}

func (cs *CaptchaService) CheckCaptchaNonce(nonceKey, nonce string) bool {
	redisKey := fmt.Sprintf(constant.REDIS_KEY_USER_CAPTCHA_NONCE_KEY, nonceKey)
	redisValue := redis.Helper.Get(redisKey)

	if strings.EqualFold(nonce, redisValue) {
		return redis.Helper.Del(redisKey)
	}

	return false
}

func (cs *CaptchaService) SaveCaptchaNonce(nonceKey, nonce string, expire time.Duration) bool {
	redisKey := fmt.Sprintf(constant.REDIS_KEY_USER_CAPTCHA_NONCE_KEY, nonceKey)
	return redis.Helper.SetWithExpir(redisKey, nonce, expire)
}

func redisCheckCaptcha(captchaId, captchaCode string) bool {
	redisKey := fmt.Sprintf(constant.REDIS_KEY_USER_CAPTCHA_KEY, captchaId)
	redisValue := redis.Helper.Get(redisKey)

	if strings.EqualFold(captchaCode, redisValue) {
		redis.Helper.Del(redisKey)
		return true
	}
	return false
}

func redisSaveCaptcha(captchaKey, captchaCode string, expire time.Duration) bool {
	redisKey := fmt.Sprintf(constant.REDIS_KEY_USER_CAPTCHA_KEY, captchaKey)
	return redis.Helper.SetWithExpir(redisKey, captchaCode, expire)
}

func getRandStr(n int) (randStr string) {
	// chars := "1234567890" // 纯数字不需要剔除
	chars := "ABCDEFGHIJKLMNPQRSTUVWXYZabcdefhjkmnpqrstxyz123456789" // 剔除了容易混淆的Oolgwvui0
	charsLen := len(chars)
	if n > 10 {
		n = 10
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		randIndex := rand.Intn(charsLen)
		randStr += chars[randIndex : randIndex+1]
	}
	return randStr
}

func imgText(width, height int, text string) (b []byte) {
	textLen := len(text)
	dc := gg.NewContext(width, height)
	bgR, bgG, bgB, bgA := getRandColorRange(240, 255)
	dc.SetRGBA255(bgR, bgG, bgB, bgA)
	dc.Clear()

	// 干扰线
	for i := 0; i < 20; i++ {
		x1, y1 := getRandPos(width, height)
		x2, y2 := getRandPos(width, height)
		r, g, b, a := getRandColor(255)
		w := float64(rand.Intn(3) + 1)
		dc.SetRGBA255(r, g, b, a)
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}

	fontSize := float64(height/2) + 5
	face := loadFontFace(fontSize)
	dc.SetFontFace(face)
	// 渲染文字
	for i := 0; i < len(text); i++ {
		r, g, b, _ := getRandColor(100)
		dc.SetRGBA255(r, g, b, 255)
		fontPosX := float64(width/textLen*i) + fontSize*0.6

		writeText(dc, text[i:i+1], float64(fontPosX), float64(height/2))
	}

	buffer := bytes.NewBuffer(nil)
	dc.EncodePNG(buffer)
	b = buffer.Bytes()
	return
}

// 渲染文字
func writeText(dc *gg.Context, text string, x, y float64) {
	xfload := 5 - rand.Float64()*10 + x
	yfload := 5 - rand.Float64()*10 + y

	radians := 40 - rand.Float64()*80
	dc.RotateAbout(gg.Radians(radians), x, y)
	dc.DrawStringAnchored(text, xfload, yfload, 0.2, 0.5)
	dc.RotateAbout(-1*gg.Radians(radians), x, y)
	dc.Stroke()
}

// 随机坐标
func getRandPos(width, height int) (x float64, y float64) {
	x = rand.Float64() * float64(width)
	y = rand.Float64() * float64(height)
	return x, y
}

// 随机颜色
func getRandColor(maxColor int) (r, g, b, a int) {
	r = int(uint8(rand.Intn(maxColor)))
	g = int(uint8(rand.Intn(maxColor)))
	b = int(uint8(rand.Intn(maxColor)))
	a = int(uint8(rand.Intn(255)))
	return r, g, b, a
}

// 随机颜色范围
func getRandColorRange(miniColor, maxColor int) (r, g, b, a int) {
	if miniColor > maxColor {
		miniColor = 0
		maxColor = 255
	}
	r = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	g = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	b = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	a = int(uint8(rand.Intn(maxColor-miniColor) + miniColor))
	return r, g, b, a
}

// 加载字体
func loadFontFace(points float64) font.Face {
	// 这里是将字体TTF文件转换成了 byte 数据保存成了一个 go 文件 文件较大可以到附录下
	// 通过truetype.Parse可以将 byte 类型的数据转换成TTF字体类型
	f, err := truetype.Parse(COMICSAN)

	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
	})
	return face
}
