package router

import (
	"github.com/CreatorQWQ/taskflow/internal/handler"
	"github.com/CreatorQWQ/taskflow/internal/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化所有路由
func InitRouter(h *handler.AllHandlers, jwtSecret string) *gin.Engine {
	r := gin.Default()

	// 跨域中间件 (可选，如果以后 Flutter 跑在 Web 端需要)
	r.Use(gin.Recovery())

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
			auth.POST("/register", h.Auth.Register)
			auth.POST("/login", h.Auth.Login)
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

		// 2. 任务路由 (受中间件保护)
		tasks := v1.Group("/tasks").Use(middleware.AuthMiddleware(jwtSecret))
		{
			tasks.POST("/", h.Task.CreateTask)              // 创建任务
			tasks.GET("/", h.Task.ListTasks)                // 获取任务列表
			tasks.PATCH("/:id/toggle", h.Task.ToggleStatus) // 使用 PATCH 表示部分更新
			tasks.DELETE("/:id", h.Task.Delete)             // 使用 DELETE 表示删除
		}

		// 以后可以在这里加任务模块路由
		// task := v1.Group("/tasks")
		// {
		//     task.POST("/", taskHandler.CreateTask)
		// }
	}

	return r
}
