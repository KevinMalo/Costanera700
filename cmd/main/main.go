package main

import (
	"github.com/kevinmalo/Costanera700/internal/logs"
)


func main()  {

	_ = logs.InitLogger()

	mux := Routes()
	server := NewServer(mux)
	server.Run()

	// Set all buyers
	//upload_buyers.SetBuyers()

	// Set all products
	//upload_products.SetBuyers()

	// Set all transactions.
	//upload_transactions.SetTransactions()
}

