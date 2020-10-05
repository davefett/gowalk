package main


// Route => "destination": ["path", "to", "get", "there"]
type Route map[string][]string

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