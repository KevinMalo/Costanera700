package upload_buyers

import (
	"encoding/json"
	"github.com/kevinmalo/Costanera700/internal/database"
	"github.com/kevinmalo/Costanera700/internal/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func SetBuyers(date int) {
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

	var buyers = []models.Buyer{}

	err = json.Unmarshal(body, &buyers)
	if err != nil {
		log.Fatal("Error al convertir a JSON: " + err.Error())
	}

	//fmt.Println(buyers[0])

	// Encode JSON

	var buyerJson = []models.Buyer{}

	for i := range buyers {
		b := models.Buyer{
					Id:   buyers[i].Id,
					Name: buyers[i].Name,
					Age:  buyers[i].Age,
					Date: date,
				}
		buyerJson = append(buyerJson,b)
	}

	data, err := json.MarshalIndent(buyerJson, "", "  ")

	if err != nil {
		log.Fatal("Error al convertir a JSON: " + err.Error())
	}

	//fmt.Printf("%s", data)

	//COMMIT
	database.Commit(data)

}