package upload_transactions

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/kevinmalo/Costanera700/internal/database"
	"github.com/kevinmalo/Costanera700/internal/models"
	"io"
	"log"
	"os"
	"regexp"
)

// Push transactions data into db
func SetTransactions(date int) {

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
	var transaction []models.Transaction

	var allIdsTransactions []string

	for {
		//Reading line by line
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading the line: %v", err)
		}

		t := models.Transaction{
			Id:         record[0],
			BuyerId:    record[1],
			Ip:         record[2],
			Device:     record[3],
			Date:     	date,
		}

		if record[4] == "" {
			log.Printf("Products ids is empty: %v", err)
			continue
		}

		//Extract products ids
		if record[4] != "" {

			//Compile regex
			re, err := regexp.Compile(`\w{7,8}`)
			if err != nil {
				log.Printf("Products ids is empty: %v", err)
			}

			//Extract ids matches
			matches := re.FindAllString(record[4],-1)
			t.ProductIds = matches

			for _, r := range matches{
				allIdsTransactions = append(allIdsTransactions, r)
			}


		}

		transaction = append(transaction, t)

	}

	// Get the best sellers
	//fmt.Println(allIdsTransactions)
	fmt.Println("******")
	printUniqueValue(allIdsTransactions)
	fmt.Println("******")

	//Create JSON
	jsonTransactions, err := json.MarshalIndent(transaction, "", "  ")
	if err != nil {
		log.Fatal("Error when encoding json: " + err.Error())
	}

	//Commit database
	database.Commit(jsonTransactions)

}

// Count the best sellers
func printUniqueValue( arr []string){
	//Create a dictionary of values for each element
	dict:= make(map[string]int)
	for _ , id :=  range arr {
		dict[id] = dict[id]+1
	}

	fmt.Println(dict)
}