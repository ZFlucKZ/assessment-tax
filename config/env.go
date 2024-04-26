package config

import (
	"os"
	"strconv"
)

type EnvMapping struct {
	Port          uint64
	DatabaseURL   string
	AdminUsername string
	AdminPassword string
}

var Env EnvMapping

func SetupEnv() {
	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 64)
	if err != nil {
		port = 8080
	}

	Env = EnvMapping{
		Port:          port,
		DatabaseURL: os.Getenv("DATABASE_URL"),
		AdminUsername: os.Getenv("ADMIN_USERNAME"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
	}
}