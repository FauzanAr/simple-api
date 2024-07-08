package config

import (
	"context"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"

	"simple-api.com/m/src/pkg/logger"
)

type Config struct {
	AppEnviroment string `env:"APP_ENVIROMENT" envDefault:"debug"`
	AppPort       string `env:"APP_PORT,required"`
	Mysql         MySql
	Jwt           Jwt
}

type MySql struct {
	Host         string `env:"MYSQL_HOST,required"`
	Port         int    `env:"MYSQL_PORT,required"`
	Password     string `env:"MYSQL_PASSWORD,required"`
	Username     string `env:"MYSQL_USERNAME,required"`
	DatabaseName string `env:"MYSQL_DATABASE_NAME,required"`
}

type Jwt struct {
	SecretKey           string `env:"JWT_SECRET_KEY,required"`
	AccessTokenExpired  int    `env:"JWT_ACCESS_TOKEN_EXPIRED_HOURS,required"`
}

func LoadEnv(ctx context.Context, log logger.Logger) (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Error(ctx, "Error while load enviroment", err, nil)
		return Config{}, err
	}

	var conf Config
	err = env.Parse(&conf)
	if err != nil {
		log.Error(ctx, "Error while parsing the enviroment", err, nil)
		fmt.Println(": ", err)
		return Config{}, err
	}

	return conf, nil
}
