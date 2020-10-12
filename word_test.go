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
	return word == "validword"
}

func (d DictionaryStub) Mutate(word string) Dictionary {
	dictionary = make(Dictionary)
	dictionary[word] = true
	return dictionary
}

func (d DictionaryStub) Add(_ string) {
	// noop
}

func TestGetWord(t *testing.T) {
	stub := &DictionaryStub{}
	wordServer := &WordServer{stub}

	router := mux.NewRouter()
	router.Handle("/words/{word:[a-zA-Z]+}", http.HandlerFunc(wordServer.WordHandler))

	t.Run("GET word returns true if valid", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/words/validword", nil)
		if err != nil {
			t.Errorf("unable to create request %v", err)
		}

		response := httptest.NewRecorder()

		router.Handle("/word/{word:[a-zA-Z]+}", http.HandlerFunc(wordServer.WordHandler))
		router.ServeHTTP(response, request)

		assertStatusCode(t, response, http.StatusOK)
		assertWordEquals(t, response, Word{"validword", true})
	})

	t.Run("GET invalid word returns false if valid", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/words/fail", nil)
		if err != nil {
			t.Errorf("unable to create request %v", err)
		}

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assertStatusCode(t, response, http.StatusOK)
		assertWordEquals(t, response, Word{"fail", false})
	})
}

func TestPutWord(t *testing.T) {
	stub := &DictionaryStub{}
	wordServer := &WordServer{stub}

	router := mux.NewRouter()
	router.Handle("/words/{word:[a-zA-Z]+}", http.HandlerFunc(wordServer.WordHandler))

	t.Run("should PUT new words into dictionary", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodPut, "/words/newword", nil)
		if err != nil {
			t.Errorf("unable to create request %v", err)
		}

		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assertStatusCode(t, response, http.StatusCreated)
	})
}

func assertStatusCode(t *testing.T, r *httptest.ResponseRecorder, expected int) {
	t.Helper()

	if actual := r.Code; actual != expected {
		t.Errorf("receive status %v, expecting %v", actual, expected)
	}
}

func assertWordEquals(t *testing.T, r *httptest.ResponseRecorder, expected Word) {
	t.Helper()

	var actual Word
	err := json.Unmarshal(r.Body.Bytes(), &actual)
	if err != nil {
		t.Errorf("unable to decode %v as json", r.Body.String())
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("received %v, expected %v", actual, expected)
	}
}
