package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Deposit struct {
	TxId          string
	Vout          int16
	Address       string
	Amount        float64
	Category      string
	Confirmations int16
}

func (d *Deposit) Save(db *sql.DB) error {
	dep := Deposit{}

	serr := db.QueryRow("SELECT txid FROM btc_deposit where txid=? and vout=?", dep.TxId, dep.Vout).Scan(&dep.TxId)

	if serr == nil {
		// record already exists, update confirmations only

		_, uerr := db.Exec("UPDATE `btc_deposit` set confirmations=? where txid=? and vout=?", d.Confirmations, d.TxId, d.Vout)

		if uerr != nil {
			log.Println("update erorr")
			log.Println(uerr)
			return fmt.Errorf("cannot update row %s", d.TxId)
		}
		return nil
	}

	if serr.Error() != "sql: no rows in result set" {
		// not standard error, return error
		log.Println(serr)
		return errors.New("backend error")
	}

	_, err := db.Exec("INSERT INTO `btc_deposit` (address, amount, category, confirmations, txid, vout) values (?, ?, ?, ?, ?, ?)", d.Address, d.Amount, d.Category, d.Confirmations, d.TxId, d.Vout)

	if err != nil {
		log.Println("insert error")
		log.Println(err, d.Address)
		return errors.New("Address, unable to insert record")
	}
	return nil
}
