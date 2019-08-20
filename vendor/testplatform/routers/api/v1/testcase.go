package v1

import (
	"net/http"
	"os"
	"log"
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"github.com/360EntSecGroup-Skylar/excelize"
	
	"testplatform/models"
	"testplatform/pkg/e"
	"testplatform/pkg/util"
)

func GetTestcase(c *gin.Context) {

	item := c.Query("item")

	pageSize := com.StrTo(c.DefaultQuery("pageSize", "10")).MustInt() // 每页显示条数
	pageCurrent := com.StrTo(c.DefaultQuery("pageCurrent", "1")).MustInt() // 当前页面
	sortOrder := c.DefaultQuery("sortOrder", "ascend") // 排序方式
	filterImportance := c.QueryArray("importance[]") // filter importance
	
	code := e.SUCCESS

	data := models.GetTestcases(item, pageSize, pageCurrent, sortOrder, filterImportance)
	totalCount := models.GetTestcasesTotal(item, filterImportance)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"results" : data,
		"totalCount" : totalCount,
	})
}

func PutTestcase(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	module := c.PostForm("module")
	importance := c.PostForm("importance")
	describe := c.PostForm("describe")
	step := c.PostForm("step")
	result := c.PostForm("result")
	item := c.PostForm("item")

	code := e.SUCCESS

	models.PutTestcases(item, id, module, importance, describe, step, result)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
	})
}

func PostTestcase(c *gin.Context) {
	preId := com.StrTo(c.PostForm("preId")).MustInt()
	module := c.PostForm("module")
	importance := c.PostForm("importance")
	describe := c.PostForm("describe")
	step := c.PostForm("step")
	result := c.PostForm("result")
	item := c.PostForm("item")

	code := e.SUCCESS

	models.PostTestcases(item, preId, module, importance, describe, step, result)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
	})
}

func DeleteTestcase(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	item := c.Query("item")

	code := e.SUCCESS

	models.DeleteTestcases(item, id)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
	})	
}

func ExportTestcase(c *gin.Context) {
	item := c.Query("item")

	data := models.GetAllTestcases(item)
	mainPath, err := os.Getwd()

	if err != nil {
		log.Println(err)
		return
	}

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx"
	path := mainPath + "/static/" + fileName
	
	util.WriteToExcel(data, path)

	defer os.Remove(path)

	c.File(path)
}

func ImportTestcase(c *gin.Context) {

	item := c.PostForm("item")
	file, _ := c.FormFile("file")

	code := e.SUCCESS

	mainPath, _ := os.Getwd()
	path := mainPath + "/static/" + file.Filename
	c.SaveUploadedFile(file, path)

	f, err := excelize.OpenFile(path)

	if err != nil {
		log.Println(err)
		return
	}

	rows, err := f.Rows("Sheet1")

	if err != nil {
		log.Println(err)
		return
	}

	defer os.Remove(path)

	// 先将原来的用例保存
	data := models.GetAllTestcases(item)
	util.WriteToExcel(data, mainPath + "/static" + "/backup/" + strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx")

	for rows.Next() {
		row, _ := rows.Columns()
		models.PutTestcases(item, com.StrTo(row[0]).MustInt(), row[1], row[2], row[3], row[4], row[5])
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
	})
}