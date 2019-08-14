package routers

import (
	"github.com/gin-gonic/gin"

	"testplatform/routers/api"
	"testplatform/routers/api/v1"
	"testplatform/pkg/setting"
	"testplatform/middleware/jwt"
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

	// 用户登录认证
	r.POST("/auth", api.PostAuth)
	// 用户注册
	r.POST("/register", api.PostRegister)
	api_v1 := r.Group("/api/v1")
	api_v1.Use(jwt.JWT())
	{
		// 获取测试用例
		api_v1.GET("/testcase", v1.GetTestcase)
		// 更新测试用例
		api_v1.PUT("/testcase", v1.PutTestcase)
		// 添加测试用例
		api_v1.POST("/testcase", v1.PostTestcase)
		// 删除用例
		api_v1.DELETE("/testcase", v1.DeleteTestcase)
		// 构造网络图数据
		api_v1.GET("/data/graph", v1.GetGraphData)
		// 上传文件到数据库
		api_v1.POST("/tool/upload", v1.PostUploadFile)
		// 转换数据格式
		api_v1.POST("/tool/format", v1.PostChangeFileFormat)
	}

	return r
}