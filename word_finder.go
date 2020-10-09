package main

import (
	"fmt"
	"syscall"
)

func main() {
	finder := NewFinder()

	route, err := finder.FindPath("charitably", "embanked")
	if err != nil {
		fmt.Errorf("encountered an error: %v", err)
		syscall.Exit(1)
	}
	fmt.Println(route)
}
