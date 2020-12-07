package main

import (
	"github.com/kevinmalo/Costanera700/internal/logs"
	"github.com/kevinmalo/Costanera700/migrations/upload_buyers"
)


func main()  {

	_ = logs.InitLogger()

	upload_buyers.SetBuyers()
}

