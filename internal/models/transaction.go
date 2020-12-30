package models

import (
	"context"
	"encoding/json"
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
		products_ids
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

type TransactionResp struct {
	Transaction []struct {
		IP      string `json:"ip"`
		BuyerID string `json:"buyer_id"`
	} `json:"transaction"`
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

func GetTransactionsIp(buyerId string) []byte {

	transactionHistoryQuery := `
	{
	  transaction(func: eq(buyer_id,%s))  {
		ip
		buyer_id
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

	var buyerIP TransactionResp
	err = json.Unmarshal(resp.Json, &buyerIP)
	if err != nil {
		log.Fatal("Error al decodificar JSON: " + err.Error())
	}

	buyersIps := GetBuyersSameIp(buyerIP.Transaction[0].IP)

	return buyersIps

}

func GetBuyersSameIp(buyerIp string) []byte {

	transactionHistoryQuery := `
	{
	  transaction(func: eq(ip,%s))  {
		ip
		buyer_id
	  }
	}
	`

	queryFormat := fmt.Sprintf(transactionHistoryQuery, buyerIp)

	dgraphClient := database.NewClient()
	txn := dgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), queryFormat)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Json

}
