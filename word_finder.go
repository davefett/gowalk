package main

import "fmt"

func main() {
	finder := NewFinder()

	route, err := finder.FindPath("charitably", "embanked")
	if err != nil {
		fmt.Errorf("encountered an error: %v", err)
	}
	fmt.Println(route)
}
