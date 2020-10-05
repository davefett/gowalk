package main

import "fmt"

func main() {
	finder := NewFinder()
	fmt.Println(finder.FindPath("charitably", "embanked"))
}
