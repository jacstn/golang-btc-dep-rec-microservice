package ext

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"

	"github.com/jacstn/golang-btc-dep-rec-microservice/internal/models"
)

func getBitcoinCliPath() string {
	return os.Getenv("BITCOIN_CLI")
}

func ListTransactions() models.RawTransactions {
	cmd := exec.Command(getBitcoinCliPath(), "listsinceblock")
	out, err := cmd.Output()

	if err != nil {
		if err.Error() == "exit status 1" {
			log.Println("bitcoin cli not running ")
		} else if err.Error() == "exit status 1" {
			log.Println("bitcoin deamon not installed or other unexpected error")
		}
		return models.RawTransactions{}
	}

	var transactions models.RawTransactions

	err = json.Unmarshal([]byte(out), &transactions)

	if err != nil {
		log.Println("Cannot parse bitcoin-cli output json")
	}
	return transactions
}
