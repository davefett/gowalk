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

// Endpoint for /routes
// TODO: move this to it's own file?
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

// Does this route contain the destination>?
func (r Route) HasDestination(dest string) bool {
	_, ok := r[dest]
	return ok
}

// Does the destination string array in the route contain the needle string?
// This is potentially one of the places to look to optimize, the array lookups, while never larger
// than 99 elements in the current implementation, are still O(n).
func (r Route) PathContains(dest string, needle string) bool {
	for _, word := range r[dest] {
		if word == needle {
			return true
		}
	}
	return false
}
