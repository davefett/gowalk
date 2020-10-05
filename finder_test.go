package main

import (
	"reflect"
	"testing"
)

func TestFinder(t *testing.T) {

	var finder *Finder = NewFinder()

	t.Run("should return error if provided with an invalid start word", func(t *testing.T) {
		//var path []string
		var err error

		_, err = finder.FindPath("asd", "bar")

		if err == nil {
			t.Errorf("Should have returned an error with an invalid start word")
		}
	})

	t.Run("should return error if provided with an invalid end word", func(t *testing.T) {
		//var path []string
		var err error

		_, err = finder.FindPath("foo", "qwe")

		if err == nil {
			t.Errorf("Should have returned an error with an invalid end word")
		}
	})

	t.Run("should return a single word if start and end are the same", func(t *testing.T) {

		path, _ := finder.FindPath("test", "test")

		if len(path) != 1 {
			t.Errorf("the path length should be 1")
		} else {
			if path[0] != "test" {
				t.Errorf("The path should contain the start word")
			}
		}

	})

	t.Run("word list should be reversed", func(t *testing.T) {
		words := []string{"one", "two", "three", "four", "five"}

		actual := reverseArray(words)

		expected := []string{"five", "four", "three", "two", "one"}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("word string not reversed corrected, got %v, expected %v", actual, expected)
		}
	})

	t.Run("make path formats path correctly", func(t *testing.T) {
		left := Route{"test": {"one", "two", "three"}}
		right := Route{"test": {"cee", "bee", "a"}}
		expected := []string{"one", "two", "three", "test", "a", "bee", "cee"}

		actual, _ := makePath(left, right, "test")

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Path string not correct: got %v, expected %v", actual, expected)
		}
	})
}
