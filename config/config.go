package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	envFilePath = "config.env"

	serverAddress = "SERVER_ADDR"
	serverPort    = "SERVER_PORT"
	dbAddress     = "DB_ADDR"
	dbPort        = "DB_PORT"
	dbName        = "DB_NAME"
	dbUser        = "DB_USER"
	dbPassword    = "DB_PASS"
)

func init() {
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Println(err)
		setDefaults()
		if err = writeConfig(envFilePath); err != nil {
			log.Fatal(err)
		}
	}
}

func ServerAddress() string {
	return os.Getenv(serverAddress)
}

func ServerPort() string {
	return os.Getenv(serverPort)
}

func DBAddress() string {
	return os.Getenv(dbAddress)
}

func DBPort() string {
	return os.Getenv(dbPort)
}

func DBName() string {
	return os.Getenv(dbName)
}

func DBUser() string {
	return os.Getenv(dbUser)
}

func DBPassword() string {
	return os.Getenv(dbPassword)
}

func setDefaults() {
	log.Println("setting defaults")
	os.Setenv(serverAddress, "127.0.0.1")
	os.Setenv(serverPort, "8080")
	os.Setenv(dbAddress, "127.0.0.1")
	os.Setenv(dbPort, "5432")
	os.Setenv(dbName, "payments")
	os.Setenv(dbUser, "payments")
	os.Setenv(dbPassword, "lth gfhjk")
}

func writeConfig(path string) error {
	log.Println("creating config.env file")
	params := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	file, err := os.OpenFile(path, params, 0600)
	if err != nil {
		return err
	}

	log.Println("saving defaults")
	_, err = fmt.Fprint(
		file,
		serverAddress+":"+ServerAddress()+"\n",
		serverPort+":"+ServerPort()+"\n",
		dbAddress+":"+DBAddress()+"\n",
		dbPort+":"+DBPort()+"\n",
		dbName+":"+DBName()+"\n",
		dbUser+":"+DBUser()+"\n",
		dbPassword+":"+DBPassword()+"\n",
	)
	if err != nil {
		return err
	}
	if err = file.Close(); err != nil {
		return err
	}
	return nil
}
