package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		ServiceName string `env:"SERVICE_NAME"`
		ServicePort string `env:"SERVICE_PORT"`
		ServiceHost string `env:"SERVICE_HOST"`

		Timezone   string `env:"SERVICE_TIMEZONE"`
		DBLogLevel string `env:"DB_LOG_LEVEL"`
		LogLevel   string `env:"LOG_LEVEL"`

		AuthorizationEnabled bool `env:"AUTHORIZATION_ENABLED"`

		Jwt        jwt
		PostgresDB rdb
		GOauth     gOauth
	}
	rdb struct {
		DBHost     string `env:"DB_HOST"`
		DBPort     string `env:"DB_PORT"`
		DBUser     string `env:"DB_USER"`
		DBPassword string `env:"DB_PASSWORD"`
		DBName     string `env:"DB_NAME"`
	}
	jwt struct {
		JWTSecret             string `env:"JWT_SECRET"`
		AccessTokenExpiration uint   `env:"ACCESS_TOKEN_EXPIRATION"`
	}
	gOauth struct {
		ClientID     string `env:"GOOGLE_OAUTH_CLIENT_ID"`
		ClientSecret string `env:"GOOGLE_OAUTH_CLIENT_SECRET"`
		RedirectURL  string `env:"GOOGLE_OAUTH_REDIRECT_URL"`
		Redirect     string `env:"GOOGLE_REDIRECT"`
	}
)

func GetConfig() Config {
	config := Config{}
	if err := cleanenv.ReadEnv(&config); err != nil {
		panic(fmt.Sprintf("Failed to load config, err: %s", err))
	}

	return config
}
