package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Environment string

const (
	EnvDev  Environment = "dev"
	EnvProd Environment = "production"
	EnvTest Environment = "test"
)

var Global Config

type (
	Config struct {
		Env  string `mapstructure:"ENV"`
		HTTP `mapstructure:",squash"`
		Log  `mapstructure:",squash"`
		PG   `mapstructure:",squash"`
		Auth `mapstructure:",squash"`
	}

	HTTP struct {
		Port int `mapstructure:"HTTP_PORT"`
	}

	Log struct {
		Level string `mapstructure:"LOG_LEVEL"`
	}

	PG struct {
		Host     string `mapstructure:"POSTGRES_HOST"`
		Port     int    `mapstructure:"POSTGRES_PORT"`
		Database string `mapstructure:"POSTGRES_DATABASE"`
		User     string `mapstructure:"POSTGRES_USER"`
		Password string `mapstructure:"POSTGRES_PASSWORD"`
		Options  string `mapstructure:"POSTGRES_OPTIONS"`
	}

	Auth struct {
		AdminUsername string `mapstructure:"ADMIN_USERNAME"`
		AdminPassword string `mapstructure:"ADMIN_PASSWORD"`
		Secret        string `mapstructure:"SECRET"`
	}
)

func Initialize(env, configFolderPath string) {
	cfg, err := Load(env, configFolderPath)
	if err != nil {
		panic(err)
	}

	Global = cfg
}

// Load reads in configurations from config files and environment variables
// It uses the following order to override configurations:
// 1. .default.env, 2. .$env.env, 3. .local.env, 4. environment variables
func Load(env, configFolderPath string) (Config, error) {
	viper.SetConfigName(".default")
	viper.SetConfigType("env")
	viper.AddConfigPath(configFolderPath)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if env != "" {
		env = strings.ToLower(env)
		viper.SetConfigName(fmt.Sprintf(".%s", env))
		viper.MergeInConfig()
	}

	viper.SetConfigName(".local")
	viper.MergeInConfig()

	viper.AutomaticEnv()

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
