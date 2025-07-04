package config

type Config struct {
	Server   ServerConfig   `yaml:"server" json:"server"`
	Storage  StorageConfig  `yaml:"storage" json:"storage"`
	Database DatabaseConfig `yaml:"database" json:"database"`
}

type ServerConfig struct {
	Port int    `yaml:"port" json:"port"`
	Mode string `yaml:"mode" json:"mode"`
}

type StorageConfig struct {
	Type      string   `yaml:"type" json:"type"`
	LocalPath string   `yaml:"local_path" json:"local_path"`
	S3        S3Config `yaml:"s3" json:"s3"`
}

type S3Config struct {
	Endpoint  string `yaml:"endpoint" json:"endpoint"`
	AccessKey string `yaml:"access_key" json:"access_key"`
	SecretKey string `yaml:"secret_key" json:"secret_key"`
	Bucket    string `yaml:"bucket" json:"bucket"`
}

type DatabaseConfig struct {
	Type string `yaml:"type" json:"type"`
	DSN  string `yaml:"dsn" json:"dsn"`
}
