package midlwre

import (
	"acttos.com/avatar/model/api/resp/rscode"
	"acttos.com/avatar/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Swagger user check, not all the requests will be accepted
func SwaggerCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.Request.Header.Get("token")
		//if token != "a5e9f901b1efc7352c38e7fc45177a22641" {
		//	c.AbortWithStatus(http.StatusNotFound)
		//	return
		//}
	}
}

func LogonCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		_, accountType, check := util.CryptoHelper.CheckJWT(token)

		if !check || accountType == "-1" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, rscode.Code(c).RSP_CODE_ACCOUNT_NOT_LOGIN_ERROR)
			return
		}
	}
}
