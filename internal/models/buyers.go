package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kevinmalo/Costanera700/internal/database"
	"log"
)

const (
	//Query for get all buyers
	buyersQuery = `
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

//Buyer model structure for map the request from db
type BuyerIdResp struct {
	Buyers []struct {
		UID  string `json:"uid"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
		Date int    `json:"date"`
	} `json:"buyers"`
}

//Buyer model structure for get buyer id
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

//Get all buyers from db
func GetBuyers() []byte {

	dgraphClient := database.NewClient()
	txn := dgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), buyersQuery)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Json

}

//Get buyer by id from db
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

//Get buyer name by id
func GetBuyerName(buyersIds []byte) []byte {

	var buyerNames BuyersIds
	err := json.Unmarshal(buyersIds, &buyerNames)
	if err != nil {
		log.Fatal("Error decoding json: " + err.Error())
	}

	var buyersNameJson = []BuyerIdResp{}

	for _, b := range buyerNames.Transaction {

		var p BuyerIdResp
		err := json.Unmarshal(GetBuyersById(b.BuyerID), &p)
		if err != nil {
			log.Fatal("Error decoding json: " + err.Error())
		}

		buyersNameJson = append(buyersNameJson,p)
	}

	data, err := json.Marshal(buyersNameJson)
	if err != nil {
		log.Fatal("Error when encoding json: " + err.Error())
	}

	return data
}
