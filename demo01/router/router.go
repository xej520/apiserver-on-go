package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xingej-go/Apiserver-go/demo01/handler/sd"
	"xingej-go/Apiserver-go/demo01/router/middleware"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	// 自定义，添加中间件
	//在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，
	//这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器
	g.Use(gin.Recovery())
	//强制浏览器不使用缓存
	g.Use(middleware.NoCache)
	g.Use(middleware.Options) // 浏览器跨域 OPTIONS 请求设置
	g.Use(middleware.Secure)  //一些安全设置
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handlers
	// sd 分组主要用来检查 API Server 的状态：
	//健康状况、服务器硬盘、CPU 和内存使用量
	// 路由分组
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
