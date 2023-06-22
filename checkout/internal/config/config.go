package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const pathToConfig = "config-dev.yaml"

type Config struct {
	HttpServer struct {
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		ShutDownTimeout int    `yaml:"shutdownTimeout"`
	} `yaml:"httpServer"`
	GrpcServer struct {
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		ShutDownTimeout int    `yaml:"shutdownTimeout"`
	} `yaml:"grpcServer"`
	Services struct {
		LOMS struct {
			Addr string `yaml:"addr"`
		} `yaml:"loms"`
		ProductService struct {
			Addr  string `yaml:"addr"`
			Token string `yaml:"token"`
			RPS   int64  `yaml:"rps"`
		} `yaml:"productService"`
	} `yaml:"services"`
	Postgres struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
		Port     int    `yaml:"port"`
	}
}

var AppConfig = Config{}

func InitConfig() error {
	bs, err := os.ReadFile(pathToConfig)
	if err != nil {
		return fmt.Errorf("[initConfig] read config file error: %w", err)
	}

	err = yaml.Unmarshal(bs, &AppConfig)
	if err != nil {
		return fmt.Errorf("[initConfig] parse bytes error: %w", err)
	}
	return nil
}

func (cfg *Config) GetHttpServerAddr() string {
	return fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port)
}

func (cfg *Config) GetGrpcServerAddr() string {
	return fmt.Sprintf("%s:%d", cfg.GrpcServer.Host, cfg.GrpcServer.Port)
}

func (cfg *Config) GetPostgresUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DB,
	)
}
