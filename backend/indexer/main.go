/*
	This package provides a pipeline to index mail data from text files
	in a local directory in a Zinc database.
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	zh "indexer/zinchandler"
	"io"
	"io/ioutil"
	"net/mail"
	"os"
	"path/filepath"
	"strings"
)

type Mail struct {
	MessageId               string
	Date                    string
	From                    string
	To                      string
	Subject                 string
	MimeVersion             string
	ContentType             string
	ContentTransferEconding string
	XFrom                   string
	XTo                     string
	XCC                     string
	XBCC                    string
	XFolder                 string
	XOrigin                 string
	XFilename               string
	Content                 string
}

func NewMailFromFile(mailData *mail.Message) (mail *Mail, err error) {
	header := mailData.Header
	content, contentErr := io.ReadAll(mailData.Body)
	if contentErr != nil {
		return nil, contentErr
	}
	return &Mail{
		MessageId:               header.Get("Message-ID"),
		Date:                    header.Get("Date"),
		From:                    header.Get("From"),
		To:                      header.Get("To"),
		Subject:                 header.Get("Subject"),
		MimeVersion:             header.Get("Mime-Version"),
		ContentType:             header.Get("Content-Type"),
		ContentTransferEconding: header.Get("Content-Transfer-Encoding"),
		XFrom:                   header.Get("X-From"),
		XTo:                     header.Get("X-To"),
		XCC:                     header.Get("X-cc"),
		XBCC:                    header.Get("X-bcc"),
		XFolder:                 header.Get("X-Folder"),
		XOrigin:                 header.Get("X-Origin"),
		XFilename:               header.Get("X-FileName"),
		Content:                 string(content),
	}, nil
}

func main() {
	var dataPath string
	flag.StringVar(&dataPath, "path", "", "Absolute path to directory")
	flag.Parse()

	if dataPath == "" {
		fmt.Println("Error: no --path argument, please provide an absolute path")
		return
	}
	err := iterateDir(dataPath)
	if err != nil {
		fmt.Printf("process failed: %s", err)
	} else {
		fmt.Println("process completed")
	}
}

func readMailFile(filePath string) (msg *mail.Message, err error) {
	file, fileErr := ioutil.ReadFile(filePath)
	if fileErr != nil {
		return nil, fmt.Errorf("error while reading file %s", fileErr)
	}
	fileContent := strings.NewReader(string(file))
	mailData, mailErr := mail.ReadMessage(fileContent)
	if mailErr != nil {
		return nil, mailErr
	}
	return mailData, nil
}

func parseFile(filePath string) (msg string, err error) {
	mailData, err := readMailFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s", err)
	}
	mail, err := NewMailFromFile(mailData)
	if err != nil {
		return "", fmt.Errorf("failed to parse email content: %s", err)
	}
	// fmt.Printf("%s\n", mail.MessageId)
	jsonData, err := json.Marshal(mail)
	if err != nil {
		return "", fmt.Errorf("failed to serialize email: %v", err)
	}
	return string(jsonData), nil
	// fmt.Printf("%+v", mail)
}

func getFileCount(path string) (int, error) {
	count := 0
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			count++
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

func progressBar(current, total int) string {
	const barLength = 50
	value := total / barLength
	done := current / value
	remaining := barLength - done
	return fmt.Sprintf(
		"[%s%s] %d/%d\n", strings.Repeat("#", done),
		strings.Repeat(" ", remaining), current, total)
}

func iterateDir(path string) error {
	fmt.Println("started indexing process")
	fileCount, err := getFileCount(path)
	if err != nil {
		return err
	}
	fmt.Printf("processing %v files\n", fileCount)
	current := 0

	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to process file %s: %w", path, err)
		}
		if info.IsDir() {
			fmt.Printf("processing directory %s\n", info.Name())
			return nil
		}
		data, err := parseFile(path)
		if err != nil {
			fmt.Printf("failed to parse file %s: %v\n", path, err)
			return nil
		}
		if err := zh.PushSingleDoc(data); err != nil {
			fmt.Printf("failed to push document: %v\n", err)
			return nil
		}
		current++
		if current%100 == 0 {
			fmt.Print(progressBar(current, fileCount))
		}
		return nil
	})
}
