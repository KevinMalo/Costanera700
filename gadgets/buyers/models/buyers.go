package models

import (
	"context"
	"errors"
	"github.com/kevinmalo/Costanera700/internal/database"
	"log"
)

var ctx = context.Background()

const (
	maxLengthInId = 8
	maxLengthInName = 100
	maxLengthInAge = 255
	q = `
	{
	  buyers(func: has(name)) {
		uid
		id
		name
		age
	  }
	}
	`
)


//Buyer model structure for buyers
type Buyer struct {
	Id   string // max 8 chars
	Name string // max 100 chars
	Age  uint8  // 0 to 255
}

//CreateBuyerCMD command to create a new review
type CreateBuyerCMD struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func (cmd *CreateBuyerCMD) validate() error {
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

	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Json

}
