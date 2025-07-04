package config

// 默认配置
var DefaultConfig = Config{
	Server: ServerConfig{
		Port: 8746,
		Mode: "debug",
	},
	Storage: StorageConfig{
		Type:      "local",
		LocalPath: "./data",
	},
	Database: DatabaseConfig{
		Type: "sqlite",
		DSN:  "cloud-for-save.db",
	},
}
