package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Database struct {
		User                  string        `yaml:"user"`
		Password              string        `yaml:"password"`
		Host                  string        `yaml:"host"`
		Name                  string        `yaml:"name"`
		MaxIdleConnections    int           `yaml:"maxIdleConnections"`
		MaxOpenConnections    int           `yaml:"maxOpenConnections"`
		MaxConnectionLifeTime time.Duration `yaml:"maxConnectionLifetime"`
		MaxConnectionIdleTime time.Duration `yaml:"maxConnectionIdletime"`
		DisableTLS            bool          `yaml:"disableTLS"`
		Debug                 bool          `yaml:"debug"`
	} `yaml:"database"`
}

// LoadAppConfig builds config for database and returns a DbConfig struct.
func LoadAppConfig() AppConfig {
	var conf AppConfig
	err := cleanenv.ReadConfig("application.yaml", &conf)
	if err != nil {
		fmt.Println(err)
	}
	return conf
}
