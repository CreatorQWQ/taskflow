package router

import (
	"github.com/CreatorQWQ/taskflow/internal/handler"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化所有路由
func InitRouter(authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	// 跨域中间件 (可选，如果以后 Flutter 跑在 Web 端需要)
	r.Use(gin.Recovery(), gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// 路由分组
	v1 := r.Group("/api/v1")
	{
		// 认证模块路由
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// 以后可以在这里加任务模块路由
		// task := v1.Group("/tasks")
		// {
		//     task.POST("/", taskHandler.CreateTask)
		// }
	}

	return r
}
