package config

import (
	"os"
)

type container struct {
	Secret string
}

var Container *container

func InitConfiguration() *container {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "7795SECRET4451"
	}
	Container = &container{
		Secret: secret,
	}
	return Container
}

func GetContainer() *container {
	return Container
}
