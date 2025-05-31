package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server     ServerConfig
	MongoDB    MongoDBConfig
	Redis      RedisConfig
	Ethereum   EthereumConfig
	JWT        JWTConfig
	AWS        AWSConfig
	Analysis   AnalysisConfig
}

type ServerConfig struct {
	Port int
	Env  string
}

type MongoDBConfig struct {
	URI      string
	Database string
}

type RedisConfig struct {
	URL      string
	Password string
}

type EthereumConfig struct {
	NodeURL  string
	ChainID  int64
}

type JWTConfig struct {
	Secret     string
	Expiration time.Duration
}

type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region         string
}

type AnalysisConfig struct {
	MaxContractSize int
	AnalysisTimeout time.Duration
}

func Load() (*Config, error) {
	port, _ := strconv.Atoi(getEnvOrDefault("SERVER_PORT", "8080"))
	chainID, _ := strconv.ParseInt(getEnvOrDefault("ETH_CHAIN_ID", "1"), 10, 64)
	maxSize, _ := strconv.Atoi(getEnvOrDefault("MAX_CONTRACT_SIZE", "500000"))
	timeout, _ := strconv.Atoi(getEnvOrDefault("ANALYSIS_TIMEOUT", "300"))

	return &Config{
		Server: ServerConfig{
			Port: port,
			Env:  getEnvOrDefault("ENV", "development"),
		},
		MongoDB: MongoDBConfig{
			URI:      getEnvOrDefault("MONGODB_URI", "mongodb://localhost:27017"),
			Database: getEnvOrDefault("MONGODB_DATABASE", "cerify"),
		},
		Redis: RedisConfig{
			URL:      getEnvOrDefault("REDIS_URL", "redis://localhost:6379"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		Ethereum: EthereumConfig{
			NodeURL:  getEnvOrDefault("ETH_NODE_URL", ""),
			ChainID:  chainID,
		},
		JWT: JWTConfig{
			Secret:     getEnvOrDefault("JWT_SECRET", ""),
			Expiration: time.Hour * 24,
		},
		AWS: AWSConfig{
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			Region:         getEnvOrDefault("AWS_REGION", "us-east-1"),
		},
		Analysis: AnalysisConfig{
			MaxContractSize: maxSize,
			AnalysisTimeout: time.Duration(timeout) * time.Second,
		},
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 