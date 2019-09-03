package v1

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"

	"testplatform/models"
	"testplatform/pkg/e"
)

func GetUser(c *gin.Context) {

	pageSize := com.StrTo(c.DefaultQuery("pageSize", "10")).MustInt() // 每页显示条数
	pageCurrent := com.StrTo(c.DefaultQuery("pageCurrent", "1")).MustInt() // 当前页面

	code := e.SUCCESS

	data := models.GetAllUsers(pageSize, pageCurrent)
	totalCount := models.GetUsersTotal()

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"results" : data,
		"totalCount" : totalCount,
	})
}