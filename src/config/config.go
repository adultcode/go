package config

import (
	"errors"
	"github.com/spf13/viper"
	//"github.com/spf13/viper/remote"
	//"github.com/spf13/viper/remote"

	"log"
	"os"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port    string
	runMode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode  bool
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	DB                 string
	minIdleConnections int
	poolSIze           int
	poolTimeout        int
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatal("Error in load conf ", err)
	}

	cfg, err := ParseConfig(v)

	if err != nil {
		log.Fatal("Error in parse conf", err)

	}
	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}

	return &cfg, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	println(filename + fileType)
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath("./config")
	v.AddConfigPath(".")
	// Add the directory where your config files are located
	v.AddConfigPath("../config") // Add another path in case the executable is run from a different location

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config2 file not found")
		}
		return nil, err
	}
	return v, nil

}

func getConfigPath(env string) string {

	if env == "docker" {
		return "../config/config-docker"
	} else if env == "production" {
		return "config/config-production"
	} else {
		return "config-development"

	}

}
