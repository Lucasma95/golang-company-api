package config

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB      DB
	Gin     Gin
	Swagger Swagger
}

type Swagger struct {
	HostName string `required:"true" split_words:"true" default:"localhost"`
}

type Gin struct {
	Mode string `required:"true" split_words:"true" default:"debug"`
}

type DB struct {
	Host                 string `required:"true" split_words:"true" default:"localhost"`
	User                 string `required:"true" split_words:"true"`
	Password             string `required:"true" split_words:"true"`
	Port                 string `required:"true" split_words:"true" default:"5432"`
	Name                 string `required:"true" split_words:"true"`
	AutomigrationEnabled bool   `split_words:"true"`
}

var once sync.Once
var config Config

func EnvVars() Config {
	once.Do(func() {
		if err := envconfig.Process("", &config); err != nil {
			panic(fmt.Sprintf("Error parsing env vars %#v", err))
		}
	})
	return config
}
