package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/spf13/viper"
)

type (
	Config struct {
		API  APIConfig      `mapstructure:"api"`
		Auth AuthConfig     `mapstructure:"auth"`
		DB   DatabaseConfig `mapstructure:"database"`
	}
	APIConfig struct {
		Port    int `mapstructure:"port"`
		Timeout int `mapstructure:"timeout"`
	}
	AuthConfig struct {
		Identifier string `mapstructure:"identifier"`
		PrivateKey string `mapstructure:"private_key"`
		Expire     string `mapstructure:"expire"`
	}
	DatabaseConfig struct {
		Postgres string `mapstructure:"postgres"`
		MySql    string `mapstructure:"mysql"`
	}
)

var (
	config *Config
)

// Init reads config from file and load to memory
func Init() error {
	en := godotenv.Load()
	if en != nil {
		log.Println(en)
	}
	v := viper.New()
	env := os.Getenv("clientname_ACCOUNT_ENV")
	fmt.Println("env ", env)
	if env == "" {
		env = "development"
	}

	v.AddConfigPath(".")
	v.SetConfigName("main")
	if err := v.ReadInConfig(); err != nil {
		log.Printf("couldn't load config: %s\n", err)
		return err
	}

	var c = config
	if err := v.Unmarshal(&c); err != nil {
		log.Printf("couldn't read config: %s\n", err)
		return err
	}
	config = c
	return nil
}

// Get return loaded config
func Get() *Config {
	return config
}

func GetConfigDir() string {
	env := os.Getenv("clientname_ACCOUNT_ENV")
	if env == "" {
		env = "local"
		return fmt.Sprintf("./configs/%s", env)

	} else if env == "development" {
		env = "development"
		return fmt.Sprintf("../../configs/%s", env)
	}
	// assume production
	return fmt.Sprintf("./../configs/%s", env)
}
