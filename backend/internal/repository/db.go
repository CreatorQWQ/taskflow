package repository

import (
	"fmt"
	"github.com/CreatorQWQ/taskflow/internal/config"
	"github.com/CreatorQWQ/taskflow/pkg/logger" 
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(cfg *config.Config) *Repository {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatalf("数据库连接失败: %v", err)
	}

	logger.Log.Info("数据库初始化成功")
	
	return &Repository{
		DB: db,
	}
}