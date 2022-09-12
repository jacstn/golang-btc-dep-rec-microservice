package handlers

import (
	"fmt"
	"net/http"

	"github.com/jacstn/golang-btc-dep-rec-microservice/config"
	"github.com/jacstn/golang-btc-dep-rec-microservice/internal/ext"
)

var app *config.AppConfig

func NewHandlers(c *config.AppConfig) {
	app = c
}

func Home(w http.ResponseWriter, r *http.Request) {
	transactions := ext.ListTransactions()
	fmt.Println(transactions)
	w.Write([]byte("ok"))
}
