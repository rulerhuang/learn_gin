package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"learn.gin/pkg/e"
	"learn.gin/pkg/util"
)

func JwtMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			code = e.SUCCESS
			data interface{}
		)

		token := ctx.Query("token")
		if token == "" {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			ctx.Abort() // 及时终止handler
		}
	}
}
