package config

import (
	"os"
)

type Config struct {
	Server *serverConfig
}

type serverConfig struct {
	Port         string
	Name         string
	Host         string
	AdminUrl      string
	ManagerUrl   string
}

func New() (*Config, error) {
	port:= os.Getenv("PORT")
	name := os.Getenv("NAME")
	host := os.Getenv("HOST")
	adminUrl := os.Getenv("ADMIN_URL")
	managerUrl := os.Getenv("MANAGER_URL")

	return &Config{
		Server: &serverConfig{
			Port:         port,
			Name:         name,
			Host:         host,
			AdminUrl:      adminUrl,
			ManagerUrl:   managerUrl,
		},
	}, nil
}
