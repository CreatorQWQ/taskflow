package main

import (
	"github.com/CreatorQWQ/taskflow/internal/config"
	"github.com/CreatorQWQ/taskflow/internal/handler"
	"github.com/CreatorQWQ/taskflow/internal/repository"
	"github.com/CreatorQWQ/taskflow/internal/router"
	"github.com/CreatorQWQ/taskflow/internal/service"
	"github.com/CreatorQWQ/taskflow/pkg/logger"
)

func main() {
	// 1. 系统初始化
	logger.InitLogger("dev")
	cfg := config.LoadConfig()

	// 2. 组装依赖 (Dependency Injection)
	repo := repository.NewRepository(cfg)
	userSvc := service.NewUserService(repo, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(userSvc)

	// 3. 初始化路由
	r := router.InitRouter(authHandler)

	logger.Log.Infof("服务正在启动，端口: %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
