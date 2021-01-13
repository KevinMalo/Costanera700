package upload_products

import (
	"encoding/csv"
	"encoding/json"
	"github.com/kevinmalo/Costanera700/internal/database"
	"github.com/kevinmalo/Costanera700/internal/models"
	"io"
	"log"
	"os"
	"strconv"
)

// Push products data into db
func SetProducts(date int) {

	//Open CSV
	f, err := os.Open("./datafiles/products/products.txt")
	if err != nil {
		log.Printf("Error while file is opening: %v", err)
	}
	defer f.Close()

	//Reading CSV
	r := csv.NewReader(f)
	r.Comma = '\''
	r.FieldsPerRecord = 3

	//Iteration CSV
	var products []models.Product
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading the line: %v", err)
		}

		c := models.Product{
			Id:   record[0],
			Name: record[1],
			Date: date,
		}

		if record[2] == "" {
			log.Printf("Price of product is empty: %v", err)
			continue
		}

		i, err := strconv.Atoi(record[2])
		if err != nil {
			log.Printf("Error reading the line: %v", err)
			continue
		}

		c.Price = i

		products = append(products, c)

	}

	//Create JSON
	jsonProduct, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatal("Error when encoding json: " + err.Error())
	}

	//Commit database
	database.Commit(jsonProduct)

}