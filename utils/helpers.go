package utils

import "go.uber.org/zap"

type Config struct {
    MongoConfig  MongoConfig  `yaml:"mongoDb"`
    ZapConfig    zap.Config   `yaml:"zap"`
    ServerConfig ServerConfig `yaml:"server"`
}

type MongoConfig struct {
    MongoUri string `yaml:"mongoUri"`
    Database string `yaml:"dataBaseName"`
}

type ServerConfig struct {
    Port string `yaml:"port"`
}