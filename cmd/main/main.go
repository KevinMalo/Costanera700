package main

import (
	"github.com/kevinmalo/Costanera700/internal/logs"
	"github.com/kevinmalo/Costanera700/migrations/upload_products"
)


func main()  {

	_ = logs.InitLogger()

	// Set all buyers
	//upload_buyers.SetBuyers

	// Set all products
	upload_products.SetBuyers()
}

