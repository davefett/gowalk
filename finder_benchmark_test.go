package main

import (
	"fmt"
	"testing"
)

func BenchmarkFinder(b *testing.B) {
	finder := NewFinder()

	output, _ := finder.FindPath("charitably", "embanked")

	fmt.Println(output)
}
