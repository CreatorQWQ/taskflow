package main

import (
	"github.com/CreatorQWQ/taskflow/internal/config"
	"github.com/CreatorQWQ/taskflow/internal/repository"
	"github.com/CreatorQWQ/taskflow/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 先初始化日志 (暂时手动传 "dev"，后面可以从环境变量读)
	logger.InitLogger("dev")
	defer logger.Log.Sync() // 程序退出前冲刷缓存区，确保日志全部写入

	// 2. 加载配置
	cfg := config.LoadConfig()

	// 3. 初始化 Repository (包含 DB 连接)
	repository.NewRepository(cfg)

	// 4. 设置 Gin (生产环境下应改为 gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		logger.Log.Info("收到 Ping 请求") // 结构化日志
		c.JSON(200, gin.H{"message": "pong"})
	})

	logger.Log.Infof("服务正在启动，端口: %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
