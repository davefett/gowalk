package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type FinderStub struct {
	// empty
}

func (fs FinderStub) FindPath(start string, end string) ([]string, error) {
	if start == "here" && end == "there" {
		return []string{"one", "two", "three"}, nil
	} else {
		return nil, fmt.Errorf("routing error from %v to %v", start, end)
	}
}

func TestGetRoute(t *testing.T) {
	t.Run("GET route", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/routes?start=here&end=there", nil)
		if err != nil {
			t.Errorf("unable to create request %v", err)
		}

		routeServer := &RouteServer{&FinderStub{}}

		router := mux.NewRouter()
		router.Path("/routes").
			HandlerFunc(routeServer.RouteHandler)

		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)
		expectedStatus := http.StatusOK

		if actual := response.Code; actual != expectedStatus {
			t.Errorf("receive status %v, expected %v", actual, expectedStatus)
		}

		var actualRoute RoutePayload
		err = json.Unmarshal(response.Body.Bytes(), &actualRoute)
		if err != nil {
			t.Errorf("error deserialize json %v.  %v", response.Body.String(), err)
		}

		expectedRoute := RoutePayload{"here", "there", []string{"one", "two", "three"}}

		if !reflect.DeepEqual(actualRoute, expectedRoute) {
			t.Errorf("received route %v, expected %v", actualRoute, expectedRoute)
		}
	})
}

func TestRoute(t *testing.T) {

	t.Run("route has a destination", func(t *testing.T) {
		route := Route{"test": []string{"test", "one", "two"}}

		if !route.HasDestination("test") {
			t.Errorf("Route should have destination %s", "test")
		}
	})

	t.Run("check if route's path contains a word", func(t *testing.T) {
		route := Route{
			"number": []string{"one", "two", "three"},
			"alpha":  []string{"ay", "bee", "cee"},
		}

		if !route.PathContains("number", "two") {
			t.Errorf("route should contain element")
		}

		if route.PathContains("alpha", "seven") {
			t.Errorf("route should not contain this element")
		}
	})

}
