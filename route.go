package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Route => "destination": ["path", "to", "get", "there"]
type Route map[string][]string

type RoutePayload struct {
	Start string   `json:"start"`
	End   string   `json:"end"`
	Path  []string `json:"path"`
}

type RouteServer struct {
	finderService FinderService
}

func (rs *RouteServer) RouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		start := r.URL.Query().Get("start")
		end := r.URL.Query().Get("end")
		log.Printf("got start: %v end: %v", start, end)
		route, err := rs.finderService.FindPath(start, end)
		if err != nil {
			fmt.Fprintf(w, "error finding route from %v to %v", start, end)
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(RoutePayload{start, end, route})
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (r Route) HasDestination(dest string) bool {
	_, ok := r[dest]
	return ok
}

func (r Route) PathContains(dest string, needle string) bool {
	for _, word := range r[dest] {
		if word == needle {
			return true
		}
	}
	return false
}
