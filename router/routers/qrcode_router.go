package routers

import (
	"acttos.com/avatar/model/api/resp"
	"acttos.com/avatar/model/api/resp/rscode"
	"encoding/base64"
	"net/http"
	"strconv"

	"acttos.com/avatar/pkg/util/logger"
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

type QrcodeRouter struct{}

// @BasePath /api/v1
// @Tags 不限用户-二维码图片相关
// @Summary 二维码图片生成接口
// @Description 根据给定的text生成二维码图片
// @Accept text/plain
// @Produce json
// @Param size query string true "需要生成的二维码的边长"
// @Param text query string true "二维码所承载的文字信息,如果是http-url,需要进行url-encode编码"
// @Success 200 {object} resp.QRCodeImg
// @Router /qrcode/gen [GET]
func (qr *QrcodeRouter) Gen(c *gin.Context) {
	size := c.Query("size")
	text := c.Query("text")

	intSize, err := strconv.Atoi(size)
	if err != nil {
		logger.Monitor.Error("The parameter size is illegal, it must be a number.")
		c.JSON(http.StatusBadRequest, rscode.Code(c).RSP_CODE_PARAM_ERROR)
		return
	}

	if len(text) == 0 {
		c.JSON(http.StatusBadRequest, rscode.Code(c).RSP_CODE_PARAM_ERROR)
		return
	}

	bytes, err := qrcode.Encode(text, qrcode.Low, intSize)
	if err != nil {
		logger.Monitor.Error("Failed to generate a QRCode png.")
		c.JSON(http.StatusBadRequest, rscode.Code(c).RSP_CODE_QRCODE_GEN_ERROR)
		return
	}

	//c.Writer.Write(bytes)
	c.JSON(http.StatusOK, resp.QRCodeImg{
		Image: "data:image/png;base64," + base64.StdEncoding.EncodeToString(bytes),
	})
}
