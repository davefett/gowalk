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
	dictionary Dictionary
}

func (ws *WordServer) WordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	word := vars["word"]

	payload := Word{word, ws.dictionary.Contains(word)}
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Fatalf("unable to serialize %v: %v", payload, err)
	}
}
