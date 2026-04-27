package config

import (
	"github.com/CreatorQWQ/taskflow/pkg/logger"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	AppEnv     string `mapstructure:"APP_ENV"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Fatalf("读取配置失败: %v", err)
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		logger.Log.Fatalf("解析配置失败: %v", err)
	}

	return &config
}
