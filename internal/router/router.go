package router

import (
	"github.com/gin-gonic/gin"
	"iot-go/internal/controller"
	"iot-go/internal/middleware"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	r := gin.New()
	// 添加 Logger 中间件
	r.Use(gin.Logger())
	// 添加 Recovery 中间件
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.AuthMiddleware())
	// 注册路由
	r.GET("/ping", controller.Ping)
	//部门路由注册
	controller.RegisterDepartmentRoutes(r.Group("/department"))
	return r
}
