package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var TextHelper *Text

func InitTextHelper() {
	once.Do(func() {
		TextHelper = &Text{}
	})
}

type Text struct{}

func (c *Text) RandomString(mustNumber bool, length int) (randStr string) {
	var chars string
	if mustNumber {
		chars = "1234567890" // 纯数字不需要剔除
	} else {
		chars = "ABCDEFGHIJKLMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz123456789" // 剔除了容易混淆的Ooli0
	}

	charsLen := len(chars)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		randIndex := rand.Intn(charsLen)
		randStr += chars[randIndex : randIndex+1]
	}
	return randStr
}

func (c *Text) RandomASCII(length int) (randASCII string) {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890~!@#$%^&*()_+=-`,./;:{}|<>]{" // 剔除了容易混淆的Ooli0
	charsLen := len(chars)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		randIndex := rand.Intn(charsLen)
		randASCII += chars[randIndex : randIndex+1]
	}
	return randASCII
}

func (c *Text) GenerateOutOrderNo(storeId, actId int64) (randStr string) {
	time := time.Now()
	// 时间格式化精确到毫秒
	nowStr := time.Format("20060102150405.000")
	nowStr = strings.ReplaceAll(nowStr, ".", "")

	return nowStr + strconv.Itoa(int(storeId)) + strconv.Itoa(int(actId))
}

func (c *Text) GenerateTransferOutBatchNo() (randStr string) {
	time := time.Now()
	// 时间格式化精确到分钟
	nowStr := time.Format("200601021504")

	// 返回精确到10分钟的字符串
	outBatchNo := nowStr[0 : len(nowStr)-1]
	return outBatchNo
}
