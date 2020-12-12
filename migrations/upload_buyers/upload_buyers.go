package upload_buyers

import (
	"github.com/kevinmalo/Costanera700/internal/database"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

	jsonBuyer := body

	//COMMIT
	database.Commit(jsonBuyer)

}

/*
QUERY
{
  buyers(func: has(name)) {
    uid
    id
    name
    age
  }
}
*/
