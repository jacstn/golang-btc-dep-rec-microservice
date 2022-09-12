package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jacstn/golang-btc-dep-rec-microservice/config"
	"github.com/jacstn/golang-btc-dep-rec-microservice/internal/database"
	"github.com/jacstn/golang-btc-dep-rec-microservice/internal/handlers"
)

const portNumber = ":3333"

var app = config.AppConfig{
	Production: false,
}

func main() {
	err := run()
	if err != nil {
		panic("error while initializing application")
	}

	handlers.NewHandlers(&app)
	fmt.Println("Starting application", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	app.DB.Close()
	if err != nil {
		log.Fatal("Cannot start server")
	}
}

func run() error {
	db := database.Connect()
	app.DB = db
	var err error

	if err != nil {
		fmt.Println("error while Reading Char Array from file")
		return err
	}

	return nil
}
