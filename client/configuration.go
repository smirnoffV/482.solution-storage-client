package client

import "github.com/caarlos0/env"

func NewConfiguration() (config Configuration) {

	if err := env.Parse(&config); err != nil {
		panic("environment configuration parsing error: " + err.Error())
	}

	return
}

type Configuration struct {
	StorageHost string `env:"STORAGE_HOST,required"`
	StoragePort string `env:"STORAGE_PORT,required"`
}
