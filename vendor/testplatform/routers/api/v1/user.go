package v1

import(
	"net/http"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"testplatform/pkg/util"	
)

func PostLogout(c *gin.Context) {
	fmt.Println("work")
	token := c.GetHeader("Authorization")
	claims, _ := util.ParseToken(token)
	claims.ExpiresAt = time.Now().Add(-1).Unix()
	c.JSON(http.StatusOK, gin.H{
		"msg" : "成功退出",
	})
}

