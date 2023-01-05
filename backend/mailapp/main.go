package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mailapp/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Mailbody struct {
	Term string `json:"term" validate:"required"`
}

type MailResponse struct {
	Documents []utils.Hits
}

func main() {
	router := chi.NewRouter()
	// r.Use(corsMiddleware)
	router.Use(middleware.Logger)
	router.Method("OPTIONS", "/api/mails", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
	}))
	router.Method("POST", "/api/mails", http.HandlerFunc(postHandler))
	address := ":3000"
	fmt.Println("Server listening on", address)
	http.ListenAndServe(address, router)
}

// func corsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		next.ServeHTTP(w, r)
// 	})
// }

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	mailSearch := Mailbody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&mailSearch)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(400), 400)
		return
	}

	result, err := utils.SearchDocs(mailSearch.Term)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if result.Hits.Total.Value == 0 {
		log.Printf("No documents found for '%s'\n", mailSearch.Term)
		http.Error(w, http.StatusText(404), 404)
		return
	}

	mailResponse := MailResponse{Documents: result.Hits.Hits}
	mailJson, err := json.Marshal(mailResponse)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(mailJson)
}
