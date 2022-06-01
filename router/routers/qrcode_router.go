package routers

import (
	"avatarmeta.cc/avatar/model/api/resp"
	"avatarmeta.cc/avatar/model/api/resp/rscode"
	"encoding/base64"
	"net/http"
	"strconv"

	"avatarmeta.cc/avatar/pkg/util/logger"
	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

type QrcodeRouter struct{}

// @BasePath /api/v1
// @Tags Utilities-QRCode Module
// @Summary QRCode generator
// @Description Generates a QRCode image with the given text
// @Accept text/plain
// @Produce json
// @Param size query string true "the length of the square holding the QRCode image"
// @Param text query string true "the text of the QRCode image. if it is a http-url, it should be encoded with 'url-encode'"
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
