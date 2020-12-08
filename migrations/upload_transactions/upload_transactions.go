package upload_transactions

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/kevinmalo/Costanera700/internal/database"
	"io"
	"log"
	"os"
	"regexp"
)

type Transaction struct {
	Id         string   `json:"id"`
	BuyerId    string   `json:"buyer_id"`
	Ip         string   `json:"ip"`
	Device     string   `json:"device"`
	ProductIds []string `json:"product_ids"`
}

var ctx = context.Background()

func SetTransactions() {

	//Open CSV
	f, err := os.Open("./datafiles/transactions/transactions.txt")
	if err != nil {
		log.Printf("error abriendo el archivo: %v", err)
	}
	defer f.Close()

	//Reading CSV
	r := csv.NewReader(f)
	r.Comma = '|'
	r.FieldsPerRecord = 5

	//Iteration CSV
	var transaction []Transaction

	for {
		//Reading line by line
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error leyendo la linea: %v", err)
		}

		t := Transaction{
			Id:         record[0],
			BuyerId:    record[1],
			Ip:         record[2],
			Device:     record[3],
		}

		if record[4] == "" {
			log.Printf("ids de productos vacios: %v", err)
			continue
		}

		//Extract products ids
		if record[4] != "" {

			//Compile regex
			re, err := regexp.Compile(`\w{7,8}`)
			if err != nil {
				log.Printf("ids de productos vacios: %v", err)
			}

			//Extract ids matches
			matches := re.FindAllString(record[4],-1)
			t.ProductIds = matches

		}

		transaction = append(transaction, t)

	}

	// Create JSON
	jsonTransactions, err := json.MarshalIndent(transaction, "", "  ")
	if err != nil {
		log.Fatal("error al convertir a JSON: " + err.Error())
	}

	fmt.Printf("%s", jsonTransactions)

	//Commit database
	Commit(jsonTransactions)

}


func Commit(p []byte) {

	//COMMIT
	dgraphClient := database.NewClient()

	mu := &api.Mutation{
		CommitNow: true,
	}

	mu.SetJson = p
	assigned, err := dgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(assigned)
}

/*
QUERY
{
  buyers(func: has(device)) {
    uid
    id
    ip
    device
	product_ids
  }
}
*/