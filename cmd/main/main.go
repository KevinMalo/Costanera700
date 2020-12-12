package main

import (
	"github.com/kevinmalo/Costanera700/internal/logs"
	"github.com/kevinmalo/Costanera700/migrations/upload_buyers"
	"github.com/kevinmalo/Costanera700/migrations/upload_products"
	"github.com/kevinmalo/Costanera700/migrations/upload_transactions"
)


func main()  {

	_ = logs.InitLogger()

	//setData()

	//Run server
	mux := Routes()
	server := NewServer(mux)
	server.Run()
}

func setData() {

	// Set all buyers
	upload_buyers.SetBuyers()

	// Set all products
	upload_products.SetBuyers()

	// Set all transactions.
	upload_transactions.SetTransactions()

}

