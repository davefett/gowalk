package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Word struct {
	Word   string `json:"word"`
	Exists bool   `json:"exists"`
}

type WordServer struct {
	dictService DictionaryService
}

// Handle the /words/ endpoint
func (ws *WordServer) WordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	if r.Method == "GET" {
		payload := Word{word, ws.dictService.Contains(word)}
		err := json.NewEncoder(w).Encode(payload)
		if err != nil {
			log.Fatalf("unable to serialize %v: %v", payload, err)
		}
	} else if r.Method == "PUT" {
		ws.dictService.Add(word)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}
