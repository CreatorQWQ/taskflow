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
	taskSvc := service.NewTaskService(repo) // 初始化任务服务
	userSvc := service.NewUserService(repo, cfg.JWTSecret)

	allHandlers := &handler.AllHandlers{
		Auth: handler.NewAuthHandler(userSvc),
		Task: handler.NewTaskHandler(taskSvc),
	}

	// 3. 初始化路由
	r := router.InitRouter(allHandlers, cfg.JWTSecret)

	logger.Log.Infof("服务正在启动，端口: %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
