package models

type RawTransaction struct {
	Address       string  `json:"address"`
	Category      string  `json:"category"`
	Amount        float64 `json:"amount"`
	Label         string  `json:"label"`
	Vout          int16   `json:"vout"`
	Confirmations int16   `json:"confirmations"`
	Blockhash     string  `json:"blockhash"`
	Blockheight   int32   `json:"blockheight"`
	Txid          string  `json:"txid"`
}

type RawTransactions struct {
	Transactions []RawTransaction `json:"transactions`
}
