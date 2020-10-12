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

	dictionaryService := NewFileDictionary("words_alpha.txt")

	wordServer := WordServer{dictionaryService}
	routeServer := RouteServer{Finder{dictionaryService}}

	router.Handle("/words/{word:[a-zA-Z]+}", http.HandlerFunc(wordServer.WordHandler)).Methods("GET")
	router.Handle("/routes", http.HandlerFunc(routeServer.RouteHandler)).Methods("PUT")

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
