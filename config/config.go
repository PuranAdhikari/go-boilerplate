package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DB  DatabaseConfig
	App AppConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

type AppConfig struct {
	Port int
}

func loadConfig(path, env string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(env) // Load based on the environment
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	return config, err
}

func InitializeConfig() Config {
	var env string
	flag.StringVar(&env, "env", "dev", "Environment (dev|prod|stg)")
	flag.Parse()
	// Load the configuration.
	cfg, err := loadConfig("./config", env)
	if err != nil {
		panic(fmt.Sprintf("Failed to load %s configuration: %v", env, err))
	}
	return cfg
}
