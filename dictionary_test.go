package main

import (
	"reflect"
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

		if !dict.Exists("test") {
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

		expected := Dictionary{
			"best":  true,
			"tet":   true,
			"tests": true,
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v to equal %v", actual, expected)
		}
	})

}
