package cmd

import (
	"net"
	"payments/config"
	"payments/db"
	"payments/repository"
	"payments/server"
	"payments/server/handlers"
	"payments/service"
)

func RunApp() (err error) {
	connDB, err := db.NewPostgresDB(db.DatabaseURL())
	if err != nil {
		return err
	}
	defer connDB.Close()

	appRepo := repository.NewRepository(connDB)
	appService := service.NewService(appRepo)
	appHandlers := handlers.NewHandlers(appService)

	listener, err := net.Listen("tcp", config.ServerPort())
	if err != nil {
		return err
	}
	defer listener.Close()

	appServer := server.NewServer(listener, appHandlers)
	if err = appServer.Run(); err != nil {
		return err
	}
	return nil
}
