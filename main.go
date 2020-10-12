package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)


var dictionary Dictionary
var finder Finder

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter()

	wordServer := WordServer{NewFileDictionary("words_alpha.txt")}
	router.Handle("/word/{word:[a-zA-Z]+}", http.HandlerFunc(wordServer.WordHandler)).Methods("GET")

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

