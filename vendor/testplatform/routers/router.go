package routers

import (
	"github.com/gin-gonic/gin"

	"testplatform/routers/api/v1"
	"testplatform/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.LoadHTMLGlob("appfront/dist/*.html")
	r.LoadHTMLFiles("appfront/dist/static/*/*")
	r.StaticFile("/", "appfront/dist/index.html")
	r.Static("/static", "appfront/dist/static")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	api_v1 := r.Group("/api/v1")
	api_v1.Use()
	{
		// 获取标签列表
		api_v1.GET("/testcase", v1.GetTestcase)
		// 构造网络图数据
		api_v1.GET("/data/graph", v1.GetGraphData)
		// 上传文件到数据库
		api_v1.POST("/tool/upload", v1.PostUploadFile)
		// 转换数据格式
		api_v1.POST("/tool/format", v1.PostChangeFileFormat)
	}

	return r
}