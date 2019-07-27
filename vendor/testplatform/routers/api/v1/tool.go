package v1

import (
	"os/exec"
	"os"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostUploadFile(c *gin.Context) {
	selectType := c.PostForm("selectType")
	if selectType == "MySQL" {
		ip := c.PostForm("ip")
		port := c.PostForm("port")
		user := c.PostForm("user")
		password := c.PostForm("password")
		dbName := c.PostForm("dbName")
		tableName := c.PostForm("tableName")
		file, _ := c.FormFile("file")
		
		mainPath, _ := os.Getwd()

		path := mainPath + "/static/" + file.Filename
		c.SaveUploadedFile(file, path)

		scriptPath := mainPath + "/vendor/testplatform/script/save_csv_to_mysql.py"

		cmd := exec.Command("python3", scriptPath, ip, port, user, password, dbName, tableName, path)
		err := cmd.Run()

		var msg string
		if err != nil {
			msg = "failed"
			fmt.Println(err)
		} else {
			msg = "success"
		}

		defer os.Remove(path)

		c.JSON(http.StatusOK, gin.H{
			"msg" : msg,
		})
	}
}


func PostChangeFileFormat (c *gin.Context) {
	targetType := c.PostForm("targetType")
	fileName := c.PostForm("fileName")
	fileType := c.PostForm("fileType")
	file, _ := c.FormFile("file")
	xsvStr := c.DefaultPostForm("xsvSplit", "x")
	sourceXsvSplit := c.DefaultPostForm("sourceXsvSplit", "x") 

	mainPath, _ := os.Getwd()

	path := mainPath + "/static/" + file.Filename
	c.SaveUploadedFile(file, path)

	scriptPath := mainPath + "/vendor/testplatform/script/change_file_format.py"
	targetPath := mainPath + "/static/" + fileName + "." + targetType 

	cmd := exec.Command("python3", scriptPath, path, fileType, targetType,targetPath, xsvStr, sourceXsvSplit)
	cmd.Run()

	c.File(targetPath)

	defer os.Remove(path)
	defer os.Remove(targetPath)
}
