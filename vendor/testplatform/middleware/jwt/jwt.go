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
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.INVALID_PARAMS
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.INVALID_PARAMS
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
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