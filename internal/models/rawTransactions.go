package models

type RawTransaction struct {
	Address       string `json:"address"`
	Category      string `json:"category"`
	Amount        string `json:"amount"`
	Label         string `json:"label"`
	Vout          string `json:"vout"`
	Confirmations string `json:"confirmations"`
	Blockhash     string `json:"blockhash"`
	Blockheight   string `json:"blockheight"`
	Txid          string `json:"txid"`
}

type RawTransactions struct {
	Transactions []RawTransaction `json:"transactions`
}
