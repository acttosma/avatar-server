package midlwre

import (
	"avatarmeta.cc/avatar/pkg/util"
	"avatarmeta.cc/avatar/pkg/util/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type bodyWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

func ParamLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		nonce := util.TextHelper.RandomString(true, 6)
		requestUrl := c.Request.RequestURI
		if strings.Contains(requestUrl, "/api/v1") {
			method := c.Request.Method
			data, err := c.GetRawData()

			if err != nil {
				logger.Monitor.Error("Error occurs when getRawData from context.", err)
			} else {
				logger.Monitor.Info("ip:", c.ClientIP(), "  user-agent:", c.Request.UserAgent(), "  content-type:", c.ContentType())
				if method != http.MethodGet {
					logger.Monitor.Infof("===== %s - Request ==========================\n%s - %s\n%s\n", nonce, method, requestUrl, string(data))
				} else {
					logger.Monitor.Infof("===== %s - Request ==========================\n%s - %s\n", nonce, method, requestUrl)
				}
			}

			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			//response写缓存
			respBodyWriter := bodyWriter{
				bodyBuf:        bytes.NewBufferString(""),
				ResponseWriter: c.Writer,
			}
			c.Writer = respBodyWriter
			c.Next()
			respBody := strings.Trim(respBodyWriter.bodyBuf.String(), "\n")
			actId := c.Request.Header.Get("accountId")
			mchId := c.Request.Header.Get("merchantId")
			logger.Monitor.Infof("===== %s - mchId:%s | actId:%s", nonce, mchId, actId)
			logger.Monitor.Infof("===== %s - Response =========================\n%s\n", nonce, respBody)
		}
	}
}
