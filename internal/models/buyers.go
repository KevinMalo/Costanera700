package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/kevinmalo/Costanera700/internal/database"
	"log"
)

const (
	maxLengthInId = 8
	maxLengthInName = 100
	maxLengthInAge = 255
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

//Buyer model structure for buyers
type Buyer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	Date int `json:"date"`
}

func (cmd *Buyer) validate() error {
	if len(cmd.Id) > maxLengthInId {
		return errors.New("id must be between 1-8 chars")
	}

	if len(cmd.Name) > maxLengthInName {
		return errors.New("name must be less than 100 chars")
	}

	if cmd.Age > maxLengthInAge {
		return errors.New("name must be less than 255 chars")
	}

	return nil
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
