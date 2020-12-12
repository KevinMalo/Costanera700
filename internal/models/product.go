package models

import (
	"context"
	"github.com/kevinmalo/Costanera700/internal/database"
	"log"
)

const productQuery = `
	{
	  products(func: has(price)) {
		uid
		id
		name
		price
		date
	  }
	}
	`

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Date int    `json:"date"`
}

func GetProducts() []byte {

	dgraphClient := database.NewClient()
	txn := dgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), productQuery)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Json

}
