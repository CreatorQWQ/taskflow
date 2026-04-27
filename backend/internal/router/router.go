package router

import (
	"github.com/CreatorQWQ/taskflow/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/CreatorQWQ/taskflow/internal/middleware"

)

// InitRouter 初始化所有路由
func InitRouter(authHandler *handler.AuthHandler, jwtSecret string) *gin.Engine {
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

		// 受保护接口：需要 Token
		// 使用我们的中间件
		user := v1.Group("/user").Use(middleware.AuthMiddleware(jwtSecret))
		{
			// 这里定义一个简单的匿名函数来测试获取个人信息
			user.GET("/me", func(c *gin.Context) {
				// 从上下文中获取中间件存入的 UserID
				userID, _ := c.Get("current_user_id")
				c.JSON(200, gin.H{
					"message": "这是私密数据",
					"user_id": userID,
				})
			})
		}

		// 以后可以在这里加任务模块路由
		// task := v1.Group("/tasks")
		// {
		//     task.POST("/", taskHandler.CreateTask)
		// }
	}

	return r
}
