package main

import (
	"testing"
)

func TestDictionary(t *testing.T) {

	t.Run("can add words to Dictionary", func(t *testing.T) {
		dict := make(Dictionary)

		dict.Add("test")

		if _, ok := dict["test"]; !ok {
			t.Errorf("word not added to Dictionary")
		}
	})

	t.Run("checking if word exists in Dictionary", func(t *testing.T) {
		dict := make(Dictionary)

		dict.Add("test")

		if !dict.Contains("test") {
			t.Errorf("exists should indicate that the word is present")
		}
	})

	t.Run("returns a list of valid word mutations", func(t *testing.T) {
		dict := make(Dictionary)
		dict.Add("test")
		dict.Add("best")
		dict.Add("tet")
		dict.Add("tests")
		dict.Add("not")
		dict.Add("negative")

		actual := dict.Mutate("test")

		expected := []string{
			"best",
			"tests",
			"tet",
		}

		if len(expected) != len(actual) {
			t.Errorf("Expected there to be %d mutations", len(expected))
		}

		for _, word := range expected {
			if !actual.Contains(word) {
				t.Errorf("Expected %s to be returned but it wasn't", word)
			}
		}
	})

}
