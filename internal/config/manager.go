package config

import (
	"cloud-for-save/pkg/logger"
	"errors"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	configFilePath = "./configs/config.yaml"
	configMu       sync.RWMutex
	currentConfig  = DefaultConfig
)

func validateConfig(cfg Config) error {
	if cfg.Server.Port <= 0 || cfg.Server.Port > 65535 {
		return errors.New("server.port 必须为 1-65535")
	}
	if cfg.Storage.Type != "local" && cfg.Storage.Type != "s3" {
		return errors.New("storage.type 仅支持 local 或 s3")
	}
	if cfg.Database.Type == "" {
		return errors.New("database.type 不能为空")
	}
	if cfg.Database.DSN == "" {
		return errors.New("database.dsn 不能为空")
	}
	return nil
}

func LoadConfig() error {
	configMu.Lock()
	defer configMu.Unlock()
	var cfg Config
	f, err := os.Open(configFilePath)
	if err == nil {
		defer f.Close()
		dec := yaml.NewDecoder(f)
		if err := dec.Decode(&cfg); err != nil {
			logger.Error("配置文件解析失败:", err)
			return err
		}
		if err := validateConfig(cfg); err != nil {
			logger.Error("配置校验失败:", err)
			return err
		}
		currentConfig = cfg
		logger.Info("已加载自定义配置文件")
		return nil
	}
	// 文件不存在或打开失败，使用默认配置
	if err := validateConfig(DefaultConfig); err != nil {
		logger.Error("默认配置校验失败:", err)
		return err
	}
	currentConfig = DefaultConfig
	logger.Info("使用默认配置")
	return nil
}

func SaveConfig(cfg Config) error {
	configMu.Lock()
	defer configMu.Unlock()
	if err := validateConfig(cfg); err != nil {
		logger.Error("配置保存校验失败:", err)
		return err
	}
	f, err := os.Create(configFilePath)
	if err != nil {
		logger.Error("配置文件写入失败:", err)
		return err
	}
	defer f.Close()
	enc := yaml.NewEncoder(f)
	defer enc.Close()
	if err := enc.Encode(&cfg); err != nil {
		logger.Error("配置文件写入失败:", err)
		return err
	}
	currentConfig = cfg
	logger.Info("配置已保存")
	return nil
}

func GetConfig() Config {
	configMu.RLock()
	defer configMu.RUnlock()
	return currentConfig
}

func ResetToDefault() error {
	return SaveConfig(DefaultConfig)
}
