package upload_products

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
	"strconv"
)

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var ctx = context.Background()

func SetBuyers() {

	//Open CSV
	f, err := os.Open("./datafiles/products/products.txt")
	if err != nil {
		log.Printf("error abriendo el archivo: %v", err)
	}
	defer f.Close()

	//Reading CSV
	r := csv.NewReader(f)
	r.Comma = '\''
	r.FieldsPerRecord = 3

	//Iteration CSV
	var products []Product
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error leyendo la linea: %v", err)
		}

		c := Product{
			Id:   record[0],
			Name: record[1],
		}

		if record[2] == "" {
			log.Printf("precio del producto vacio: %v", err)
			continue
		}

		i, err := strconv.Atoi(record[2])
		if err != nil {
			log.Printf("error leyendo la linea: %v", err)
			continue
		}

		c.Price = i

		products = append(products, c)

	}

	//Create JSON
	jsonProduct, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatal("error al convertir a JSON: " + err.Error())
	}

	fmt.Printf("%s", jsonProduct)

	//Commit database
	Commit(jsonProduct)

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
  buyers(func: has(price)) {
    uid
    id
    name
    price
  }
}
*/