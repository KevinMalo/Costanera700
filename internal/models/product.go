package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kevinmalo/Costanera700/internal/database"
	"log"
)

//Query for get all products
const productQuery = `
	{
	  products(func: has(price)) {
		uid
		product_id
		name
		price
		date
	  }
	}
	`

//Product model structure for buyers
type Product struct {
	Id    string `json:"product_id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Date  int    `json:"date"`
}

//ProductResp model structure for map the request from db
type ProductResp struct {
	Product []struct {
		UID       string `json:"uid"`
		ProductID string `json:"product_id"`
		Name      string `json:"name"`
		Price     int    `json:"price"`
		Date      int    `json:"date"`
	} `json:"product"`
}

//ProductIdResp model structure for map the request from db
type ProductIdResp struct {
	Transaction []struct {
		ProductsIds []string `json:"products_ids"`
	} `json:"transaction"`
}

//Get all products from db
func GetProducts() []byte {

	dgraphClient := database.NewClient()
	txn := dgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), productQuery)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Json

}

//Get product from db by id
func GetProductById(productId string) []byte {

	productByIdQuery := `
	{
	  product(func: eq(product_id,%s), first: 1) {
		uid
		product_id
		name
		price
		date
	  }
	}
	`
	queryFormat := fmt.Sprintf(productByIdQuery, productId)

	dgraphClient := database.NewClient()
	txn := dgraphClient.NewTxn()

	resp, err := txn.Query(context.Background(), queryFormat)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", resp.Json)
	return resp.Json

}

//Get product names from db
func GetProductsNames(productsIds []byte) []byte {

	var productsNames ProductIdResp
	err := json.Unmarshal(productsIds, &productsNames)
	if err != nil {
		log.Fatal("Error decoding json: " + err.Error())
	}

	var productJson = []ProductResp{}

	for _, t := range productsNames.Transaction {
		for _, id := range t.ProductsIds {

			var p ProductResp
			err := json.Unmarshal(GetProductById(id), &p)
			if err != nil {
				log.Fatal("Error decoding json: " + err.Error())
			}

			productJson = append(productJson, p)
		}
	}

	data, err := json.Marshal(productJson)
	if err != nil {
		log.Fatal("Error when encoding json: " + err.Error())
	}

	return data

}

//Get most selled products
func GetBestSellers() []byte {

	productsIds := [5]string{"3d659163", "7eeb79ef", "7fbe369", "979ed1c3", "d6e2c22d"}

	var productJson = []ProductResp{}

	for _, id := range productsIds {

		var p ProductResp
		err := json.Unmarshal(GetProductById(id), &p)
		if err != nil {
			log.Fatal("Error decoding json: " + err.Error())
		}

		productJson = append(productJson, p)
	}

	data, err := json.Marshal(productJson)
	if err != nil {
		log.Fatal("Error when encoding json: " + err.Error())
	}

	return data

}
