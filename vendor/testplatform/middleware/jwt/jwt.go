package jwt

import (
	"time"
	"net/http"
	
	"github.com/gin-gonic/gin"
	
	"testplatform/pkg/util"
	"testplatform/pkg/e"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 超时暂时不处理
				// code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				claims.ExpiresAt += time.Now().Add(1 * time.Hour).Unix()
			} else {
				// 每次向后端的有效请求，都会刷新token的过期时间
				claims.ExpiresAt += time.Now().Add(1 * time.Hour).Unix()
			}
			
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}