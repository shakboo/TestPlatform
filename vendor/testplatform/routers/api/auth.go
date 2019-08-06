package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	
	"testplatform/pkg/util"
	"testplatform/models"
	"testplatform/pkg/e"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func PostAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH_TOKEN
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})

	
}

func PostRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	if ok {
		isRegister := models.AddAuth(username, password)
		if isRegister {
			c.JSON(http.StatusOK, gin.H{
				"msg": "注册成功",
			})
		} else {

		}
	}
}