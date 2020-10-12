package main

import (
	"fmt"
	"testing"
)

func BenchmarkFinder(b *testing.B) {
	finder := Finder{NewFileDictionary("words_alpha.txt")}

	output, _ := finder.FindPath("charitably", "embanked")

	fmt.Println(output)
}
