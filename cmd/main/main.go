package main

import (
	"github.com/kevinmalo/Costanera700/internal/logs"
	"github.com/kevinmalo/Costanera700/migrations/upload_transactions"
)


func main()  {

	_ = logs.InitLogger()

	// Set all buyers
	//upload_buyers.SetBuyers()

	// Set all products
	//upload_products.SetBuyers()

	// Set all transactions.
	upload_transactions.SetTransactions()

	//Run server
	mux := Routes()
	server := NewServer(mux)
	server.Run()
}

