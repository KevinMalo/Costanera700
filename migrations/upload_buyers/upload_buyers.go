package upload_buyers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/kevinmalo/Costanera700/internal/database"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var ctx = context.Background()

//Buyer Defino los tipos de Datos de mi aplicación.
type Buyer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func SetBuyers()  {
	url := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Costanera-700")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var buyers = []Buyer{}

	err = json.Unmarshal(body, &buyers)
	if err != nil {
		log.Fatal("Error al convertir a JSON: " + err.Error())
	}

	//return body
	dgraphClient := database.NewClient()

	mu := &api.Mutation{
		CommitNow: true,
	}

	mu.SetJson = body
	assigned, err := dgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(assigned)
}

//func SetBuyers() {
//
//	pb := GetBuyersData()
//
//	dgraphClient := database.NewClient()
//
//	mu := &api.Mutation{
//		CommitNow: true,
//	}
//
//	mu.SetJson = pb
//	assigned, err := dgraphClient.NewTxn().Mutate(ctx, mu)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(assigned)
//
//}
