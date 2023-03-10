package zinchandler

import (
	"fmt"
	cons "indexer/constants"
	"io"
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

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func CreateNewIndex() error {
	dbUrl := os.Getenv("ZINC_DB_URL")
	url := fmt.Sprintf("%s/api/index", dbUrl)
	req, err := http.NewRequest("POST", url, strings.NewReader(cons.IndexSchema))
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(`failed to create new index: %s`, string(body))
	}
	fmt.Println(string(body))
	return nil
}
