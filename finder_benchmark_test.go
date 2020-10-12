package main

import (
	"fmt"
	"testing"
)

func BenchmarkFinder(b *testing.B) {
	finder := Finder{NewDictionary("words_alpha.txt")}

	output, _ := finder.FindPath("charitably", "embanked")

	fmt.Println(output)
}
