package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kevinmalo/Costanera700/internal/database"
	"log"
)

const (
	buyersQuery     = `
	{
	  buyers(func: has(name)) {
		uid
		id
		name
		age
		date
	  }
	}
	`
)

type TransactionName struct {
	Transaction []struct {
		IP      string `json:"ip"`
		BuyerID string `json:"buyer_id"`
	} `json:"transaction"`
}

type BuyerIdResp struct {
	Buyers []struct {
		UID  string `json:"uid"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
		Date int    `json:"date"`
	} `json:"buyers"`
}

type BuyersIds struct {
	Transaction []struct {
		IP      string `json:"ip"`
		BuyerID string `json:"buyer_id"`
	} `json:"transaction"`
}

//Buyer model structure for buyers
type Buyer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	Date int    `json:"date"`
}

func GetBuyers() []byte {

	dgraphClient := database.NewClient()
	txn := dgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), buyersQuery)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Json

}

func GetBuyersById(buyerId string) []byte {

	buyerByIdQuery := `
	{
	  buyers(func: eq(id,%s)) {
		uid
		id
		name
		age
		date
	  }
	}
	`
	queryFormat := fmt.Sprintf(buyerByIdQuery, buyerId)

	dgraphClient := database.NewClient()
	txn := dgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), queryFormat)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Json

}

func GetBuyerName(buyersids []byte) []byte {

	var buyerNames BuyersIds
	err := json.Unmarshal(buyersids, &buyerNames)
	if err != nil {
		log.Fatal("Error al decodificar JSON: " + err.Error())
	}

	var productJson = []BuyerIdResp{}

	for _, b := range buyerNames.Transaction {

		var p BuyerIdResp
		err := json.Unmarshal(GetBuyersById(b.BuyerID), &p)
		if err != nil {
			log.Fatal("Error al decodificar JSON: " + err.Error())
		}

		productJson = append(productJson,p)
	}

	data, err := json.Marshal(productJson)
	if err != nil {
		log.Fatal("Error al convertir a JSON: " + err.Error())
	}

	fmt.Printf("%s", data)

	return data
}
