package config

import (
	"github.com/caarlos0/env/v9"
)

// Config Configration Struct for config.ini
type Config struct {
	Port        int    `env:"PORT" envDefault:"8080"`
	SteamAPIKey string `env:"STEAM_API_KEY"`
	DefaultPage string `env:"DEFAULT_PAGE"`

	// Database writing
	DBWriteHost string `env:"DB_WRITE_HOST,required"`
	DBWritePort int    `env:"DB_WRITE_PORT" envDefault:"3306,required"`
	DBWriteUser string `env:"DB_WRITE_USER,required"`
	DBWritePass string `env:"DB_WRITE_PASS,unset,required"`
	DBWriteName string `env:"DB_WRITE_NAME,required"`

	// Database reading
	DBReadHost string `env:"DB_READ_HOST"`
	DBReadPort int    `env:"DB_READ_PORT" envDefault:"3306"`
	DBReadUser string `env:"DB_READ_USER"`
	DBReadPass string `env:"DB_READ_PASS,unset"`
	DBReadName string `env:"DB_READ_NAME"`

	// UserMaxResources UserMaxResources
	SecretMey string `env:"SECRET_KEY,unset,required"`
}

/*
type UserMaxResources struct {
	Servers uint16
	Teams   uint16
	Matches uint16
}
*/

func GetConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	return cfg
}
