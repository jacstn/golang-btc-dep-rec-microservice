package handlers

import (
	"net/http"

	"github.com/jacstn/golang-btc-dep-rec-microservice/config"
	"github.com/jacstn/golang-btc-dep-rec-microservice/internal/ext"
	"github.com/jacstn/golang-btc-dep-rec-microservice/internal/models"
)

var app *config.AppConfig

func NewHandlers(c *config.AppConfig) {
	app = c
}

func Home(w http.ResponseWriter, r *http.Request) {
	transactions, err := ext.ListTransactions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for _, rt := range transactions.Transactions {
		var t models.Deposit
		t.Address = rt.Address
		t.Amount = rt.Amount
		t.Category = rt.Category
		t.TxId = rt.Txid
		t.Vout = rt.Vout
		t.Confirmations = rt.Confirmations

		err = t.Save(app.DB)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.Write([]byte("ok"))
}
