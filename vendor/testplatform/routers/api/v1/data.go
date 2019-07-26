package v1

import (
	"os/exec"
	"time"
	"strconv"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetGraphData(c *gin.Context) {
	graphType := c.Query("graphType")
	nodeNum := c.Query("node")

	mainPath, _ := os.Getwd()

	timeUnix := time.Now().Unix()
	path := mainPath + "/static/" + graphType + "_" + strconv.FormatInt(timeUnix, 10) + ".csv"

	// 调用python脚本
	var scriptPath string
	if graphType == "weight" {
		scriptPath = mainPath + "/vendor/testplatform/script/weight_graph.py"
	} else {
		scriptPath = mainPath + "/vendor/testplatform/script/unweight_graph.py"
	}

	args := []string{scriptPath, nodeNum, path}
	cmd := exec.Command("python3", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	// 返回文件流
	c.File(path)
	// 删除文件
	defer os.Remove(path)

}