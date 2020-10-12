package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	router.Path("/words/{word:[a-zA-Z]+}").
		HandlerFunc(wordServer.WordHandler)
	router.Path("/routes").
		HandlerFunc(routeServer.RouteHandler)

	serverAddress := "127.0.0.1:8080"

	server := &http.Server{
		Handler:      router,
		Addr:         serverAddress,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	go func() {
		log.Printf("Starting server at %v\n", serverAddress)
		err := server.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	signalChannel := make(chan os.Signal, 1)

	// block until we receive a signal
	<-signalChannel

	signal.Notify(signalChannel, os.Interrupt)

	log.Println("received shutdown signal...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("unable to shutdown server cleanly, exiting anyway: %v\n", err)
	}
	log.Println("shut down successfully")

	// Signal clean exit
	os.Exit(0)
}
