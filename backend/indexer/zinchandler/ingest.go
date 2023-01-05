package zinchandler

import (
	"fmt"
	cons "indexer/constants"
	"net/http"
	"os"
	"strings"
)

func PushSingleDoc(data string) error {
	dbUrl := os.Getenv("ZINC_DB_URL")
	url := fmt.Sprintf("%s/api/%s/_doc", dbUrl, cons.DatabaseIndex)
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return err
	}
	dbUser := os.Getenv("ZINC_DB_USER")
	dbPassword := os.Getenv("ZINC_DB_PASSWORD")
	req.SetBasicAuth(dbUser, dbPassword)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", cons.UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// log.Println(resp.StatusCode)
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(body))
	return nil
}

// func createNewIndex() {
// 	index := `
// 	{
// 		"name": "mail",
// 		"storage_type": "disk",
// 		"shard_num": 1,
// 		"mappings": {
// 			"properties": {
// 				"Content": {
// 					"type": "text",
// 					"store": true,
// 					"highlightable": false
// 				}
// 		}
// 	}
// 	`
// 	fmt.Println(index)
// }
