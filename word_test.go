package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestWord(t *testing.T) {

	t.Run("GET word returns true if valid", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/word/test", nil)
		if err != nil {
			t.Errorf("unable to create request %v", err)
		}

		response := httptest.NewRecorder()

		wordServer := WordServer{NewDictionary("words_alpha.txt")}
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
}
