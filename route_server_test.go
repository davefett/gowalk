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
