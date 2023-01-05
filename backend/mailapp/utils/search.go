package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	cons "mailapp/constants"
	"net/http"
	"os"
	"strings"
	"time"
)

type ZincResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []Hits  `json:"hits"`
	} `json:"hits"`
}

type Hits struct {
	Index     string    `json:"_index"`
	Type      string    `json:"_type"`
	ID        string    `json:"_id"`
	Score     float64   `json:"_score"`
	Timestamp time.Time `json:"@timestamp"`
	Source    Source    `json:"_source"`
}

type Source struct {
	Content                 string `json:"Content"`
	ContentTransferEconding string `json:"ContentTransferEconding"`
	ContentType             string `json:"ContentType"`
	Date                    string `json:"Date"`
	From                    string `json:"From"`
	MessageID               string `json:"MessageId"`
	MimeVersion             string `json:"MimeVersion"`
	Subject                 string `json:"Subject"`
	To                      string `json:"To"`
	Xbcc                    string `json:"XBCC"`
	Xcc                     string `json:"XCC"`
	XFilename               string `json:"XFilename"`
	XFolder                 string `json:"XFolder"`
	XFrom                   string `json:"XFrom"`
	XOrigin                 string `json:"XOrigin"`
	XTo                     string `json:"XTo"`
}

func SearchDocs(term string) (ZincResponse, error) {
	dbUrl := os.Getenv("ZINC_DB_URL")
	query := fmt.Sprintf(cons.QueryString, term, cons.ResultsFrom, cons.ResultsMax)

	req, err := http.NewRequest(
		"POST", fmt.Sprintf("%s/api/mails/_search", dbUrl), strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}

	dbUser := os.Getenv("ZINC_DB_USER")
	dbPasword := os.Getenv("ZINC_DB_PASSWORD")
	req.SetBasicAuth(dbUser, dbPasword)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", cons.UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%s", string(body))
	var response ZincResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf(err.Error())
	}
	return response, nil
}
