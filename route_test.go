package main

import (
	"testing"
)

func TestRoute(t *testing.T) {

	t.Run("route has a destination", func(t *testing.T) {
		route := Route{"test": []string{"test", "one", "two"}}

		if !route.HasDestination("test") {
			t.Errorf("Route should have destination %s", "test")
		}
	})

	t.Run("check if route's path contains a word", func(t *testing.T) {
		route := Route{
			"number": []string{"one", "two", "three"},
			"alpha":  []string{"ay", "bee", "cee"},
		}

		if !route.PathContains("number", "two") {
			t.Errorf("route should contain element")
		}

		if route.PathContains("alpha", "seven") {
			t.Errorf("route should not contain this element")
		}
	})

}
