package config

import (
	"github.com/caarlos0/env/v9"
)

var (
	mapPool = []string{
		"de_inferno",
		"de_mirage",
		"de_nuke",
		"de_overpass",
		"de_vertigo",
		"de_ancient",
		"de_anubis",
	}
)

// Config Configration Struct for config.ini
type Config struct {
	Port        int    `env:"PORT" envDefault:"8080"`
	SteamAPIKey string `env:"STEAM_API_KEY"`
	DefaultPage string `env:"DEFAULT_PAGE"`

	// Database writing
	DBWriteHost string `env:"DB_WRITE_HOST"`
	DBWritePort int    `env:"DB_WRITE_PORT" envDefault:"3306"`
	DBWriteUser string `env:"DB_WRITE_USER"`
	DBWritePass string `env:"DB_WRITE_PASS,unset"`
	DBWriteName string `env:"DB_WRITE_NAME"`

	// Database reading
	DBReadHost string `env:"DB_READ_HOST"`
	DBReadPort int    `env:"DB_READ_PORT" envDefault:"3306"`
	DBReadUser string `env:"DB_READ_USER"`
	DBReadPass string `env:"DB_READ_PASS,unset"`
	DBReadName string `env:"DB_READ_NAME"`

	Cookie string `env:"COOKIE"`
	// UserMaxResources UserMaxResources
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
