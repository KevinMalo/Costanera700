package models

import (
	"context"
	"fmt"
	"github.com/kevinmalo/Costanera700/internal/database"
	"log"
)

const (
	transactionQuery = `
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
)

type Transaction struct {
	Id         string   `json:"id"`
	BuyerId    string   `json:"buyer_id"`
	Ip         string   `json:"ip"`
	Device     string   `json:"device"`
	Date       int      `json:"date"`
	ProductIds []string `json:"products_ids"`
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

func GetTransactionsHistory(buyerId string) []byte {

	transactionHistoryQuery := `
	{
	  transaction(func: eq(buyer_id,%s))  {
		products_ids
	  }
	}
	`

	queryFormat := fmt.Sprintf(transactionHistoryQuery, buyerId)

	dgraphClient := database.NewClient()
	txn := dgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), queryFormat)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Json

}
