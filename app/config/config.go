package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	URL     string
	Timeout int
}

type JWTConfig struct {
	RefreshTokenJwt string
	AccessTokenJwt  string
}

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
}

var config *Config

func (c *Config) SetConfig() *Config {
	databaseURL := os.Getenv("DATABASE_CONNECTED")
	databaseTimeoutStr := os.Getenv("DATABASE_TIMEOUT")
	databaseTimeout, err := strconv.Atoi(databaseTimeoutStr)
	if err != nil {
		databaseTimeout = 30000
	}

	jwtRefreshTokenJwt := os.Getenv("REFRESH_TOKEN_JWT")
	jwtAccessTokenJwt := os.Getenv("ACCESS_TOKEN_JWT")

	config = &Config{
		Database: DatabaseConfig{
			URL:     databaseURL,
			Timeout: databaseTimeout,
		},
		JWT: JWTConfig{
			RefreshTokenJwt: jwtRefreshTokenJwt,
			AccessTokenJwt:  jwtAccessTokenJwt,
		},
	}
	return config
}

func GetConfig() *Config {
	return config
}

func NewConfig(conf string) error {
	err := godotenv.Load(conf)
	if err != nil {
		return errors.New("error load config")
	}
	config.SetConfig()
	return nil
}

//func SetConfig() *Config {
//	conf := &Config{}
//	conf.Database.URL =
//	conf.Database.Timeout = 30000
//	return conf
//}

//func GetConfig() *Config {
//	return config
//}
