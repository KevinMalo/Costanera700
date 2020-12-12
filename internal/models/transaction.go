package models

import (
	"context"
	"github.com/kevinmalo/Costanera700/internal/database"
	"log"
)

const transactionQuery = `
	{
	  transaction(func: has(device)) {
		uid
		id
		ip
		device
		date
		product_ids
	  }
	}
	`

type Transaction struct {
	Id         string   `json:"id"`
	BuyerId    string   `json:"buyer_id"`
	Ip         string   `json:"ip"`
	Device     string   `json:"device"`
	Date       int      `json:"date"`
	ProductIds []string `json:"product_ids"`
}

func GetTransactions() []byte {

	dgraphClient := database.NewClient()
	txn := dgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), transactionQuery)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Json

}
