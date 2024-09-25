package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
	"log"
	"strings"
)

type Config struct {
	Server ServerConfig
	Db     SqliteConfig
	Log    LogConfig
}

type SqliteConfig struct {
	Path string
}

type LogConfig struct {
	Level string
}

func (c LogConfig) GetLevel() logger.LogLevel {
	switch strings.ToLower(c.Level) {
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	default:
		return logger.Info
	}
}

type ServerConfig struct {
	Port int
}

var AppConfig *Config

// InitConfig 初始化配置
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	AppConfig = &Config{}

	err := viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %s", err)
	}
}
