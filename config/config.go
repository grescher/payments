package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}
}

func ServerAddress() string {
	return os.Getenv("SERVER_ADDR")
}

func ServerPort() string {
	return os.Getenv("SERVER_PORT")
}

func DBAddr() string {
	return os.Getenv("DB_ADDR")
}

func DBPort() string {
	return os.Getenv("DB_PORT")
}

func DBName() string {
	return os.Getenv("DB_NAME")
}

func DBUser() string {
	return os.Getenv("DB_USER")
}

func DBPassword() string {
	return os.Getenv("DB_PASS")
}
