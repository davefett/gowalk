package main

// Route => "destination": ["path", "to", "get", "there"]
type Route map[string][]string

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
