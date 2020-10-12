package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type DictionaryStub struct {
}

func (d DictionaryStub) Contains(word string) bool {
	return word == "test"
}

func (d DictionaryStub) Mutate(word string) Dictionary {
	dictionary = make(Dictionary)
	dictionary[word] = true
	return dictionary
}

func TestWord(t *testing.T) {

	t.Run("GET word returns true if valid", func(t *testing.T) {
		var stub DictionaryService
		stub = DictionaryStub{}
		wordServer := WordServer{stub}

		request, err := http.NewRequest(http.MethodGet, "/word/test", nil)
		if err != nil {
			t.Errorf("unable to create request %v", err)
		}

		response := httptest.NewRecorder()

		router := mux.NewRouter()

		router.Handle("/word/{word:[a-zA-Z]+}", http.HandlerFunc(wordServer.WordHandler))
		router.ServeHTTP(response, request)

		if status := response.Code; status != http.StatusOK {
			t.Errorf("receive status %v, expecting %v", status, 200)
		}

		expected := Word{"test", true}
		var actual Word
		err = json.Unmarshal(response.Body.Bytes(), &actual)
		if err != nil {
			t.Errorf("unable to decode %v as json", response.Body.String())
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("received %v, expected %v", actual, expected)
		}
	})

	t.Run("GET invalid word returns false if valid", func(t *testing.T) {
		var stub DictionaryService
		stub = DictionaryStub{}
		wordServer := WordServer{stub}

		request, err := http.NewRequest(http.MethodGet, "/word/fail", nil)
		if err != nil {
			t.Errorf("unable to create request %v", err)
		}

		response := httptest.NewRecorder()

		router := mux.NewRouter()

		router.Handle("/word/{word:[a-zA-Z]+}", http.HandlerFunc(wordServer.WordHandler))
		router.ServeHTTP(response, request)

		if status := response.Code; status != http.StatusOK {
			t.Errorf("receive status %v, expecting %v", status, 200)
		}

		expected := Word{"fail", false}
		var actual Word
		err = json.Unmarshal(response.Body.Bytes(), &actual)
		if err != nil {
			t.Errorf("unable to decode %v as json", response.Body.String())
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("received %v, expected %v", actual, expected)
		}
	})
}
