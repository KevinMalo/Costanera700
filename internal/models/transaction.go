package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kevinmalo/Costanera700/internal/database"
	"log"
)

const (
	//Query for get all transactions
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

//Transaction model structure for transactions
type Transaction struct {
	Id         string   `json:"id"`
	BuyerId    string   `json:"buyer_id"`
	Ip         string   `json:"ip"`
	Device     string   `json:"device"`
	Date       int      `json:"date"`
	ProductIds []string `json:"products_ids"`
}

//Transaction model structure for map the request from db
type TransactionResp struct {
	Transaction []struct {
		IP      string `json:"ip"`
		BuyerID string `json:"buyer_id"`
	} `json:"transaction"`
}

//Get all transactions from db
func GetTransactions() []byte {

	dgraphClient := database.NewClient()
	txn := dgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), transactionQuery)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Json

}

//Get all transaction history from db for buyer
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

//Get ip buyer from transaction by id
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
		log.Fatal("Error decoding json: " + err.Error())
	}

	buyersIps := GetBuyersSameIp(buyerIP.Transaction[0].IP)

	return buyersIps

}

//Get buyers with same ip from db
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
