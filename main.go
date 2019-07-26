package main

import (
	"net/http"
	"fmt"

	"testplatform/routers"
	"testplatform/pkg/setting"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%d", setting.HTTPPort),   // 监听的TCP地址，格式为:8000
		Handler:           router,                                 // http句柄，用于处理程序响应HTTP请求
		ReadTimeout:       setting.ReadTimeout,                    // 允许读取的最大时间
		WriteTimeout:      setting.WriteTimeout,                   // 允许写入的最大时间
		MaxHeaderBytes:    1 << 20,                                // 请求头的最大字节数
		// TLSConfig  *tls.Config   安全传输协议的配置
		// IdleTimeout  time.Duration  等待的最大时间
		// ConnState    func(net.Conn, Connstate)   指定一个可选的回调函数，当客户端连接发生变化时调用
		// ErrorLog   *log.Logger   指定一个可选的日志记录器，用于接收程序的意外行为和底层系统错误。如果是nil，则在控制台输出
	}

	s.ListenAndServe()
}