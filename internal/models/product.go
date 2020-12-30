package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kevinmalo/Costanera700/internal/database"
	"log"
)

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

type Product struct {
	Id    string `json:"product_id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Date  int    `json:"date"`
}

type ProductResp struct {
	Product []struct {
		UID       string `json:"uid"`
		ProductID string `json:"product_id"`
		Name      string `json:"name"`
		Price     int    `json:"price"`
		Date      int    `json:"date"`
	} `json:"product"`
}

type MyJsonName struct {
	Transaction []struct {
		ProductsIds []string `json:"products_ids"`
	} `json:"transaction"`
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

func GetProductsNames(productsIds []byte) []byte {
	//s := string(productsIds)
	//fmt.Printf(s)

	var productsNames MyJsonName
	err := json.Unmarshal(productsIds, &productsNames)
	if err != nil {
		log.Fatal("Error al decodificar JSON: " + err.Error())
	}

	var productJson = []ProductResp{}

	for _, t := range productsNames.Transaction {
		for _, id := range t.ProductsIds {

			var p ProductResp
			err := json.Unmarshal(GetProductById(id), &p)
			if err != nil {
				log.Fatal("Error al decodificar JSON: " + err.Error())
			}

			productJson = append(productJson, p)
		}
	}

	data, err := json.Marshal(productJson)
	if err != nil {
		log.Fatal("Error al convertir a JSON: " + err.Error())
	}

	//fmt.Printf("%s", data)

	return data

}

func GetBestSellers() []byte {
	//s := string(productsIds)
	//fmt.Printf(s)

	productsIds := [5]string{"3d659163", "7eeb79ef", "7fbe369", "979ed1c3", "d6e2c22d"}

	var productJson = []ProductResp{}

	for _, id := range productsIds {

		var p ProductResp
		err := json.Unmarshal(GetProductById(id), &p)
		if err != nil {
			log.Fatal("Error al decodificar JSON: " + err.Error())
		}

		productJson = append(productJson, p)
	}

	data, err := json.Marshal(productJson)
	if err != nil {
		log.Fatal("Error al convertir a JSON: " + err.Error())
	}

	//fmt.Printf("%s", data)

	return data

}
