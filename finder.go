package main

import (
	"fmt"
)

type FinderService interface {
	FindPath(start string, end string) ([]string, error)
}

type Finder struct {
	dict DictionaryService
}

// Route: map("key") path []string

// Find a path between two strings by computing all possible permutations of a word in an array,
// and seeing if it shares that word with the words in the other 'route'.  The start and end words
// must be present in the dictionary.  If a match is not found, the function will 'extend' the
// smaller of the two routes and recursively call itself with the new route.
func (f Finder) FindPath(start string, end string) ([]string, error) {

	if !f.dict.Contains(start) {
		err := fmt.Errorf("starting word %s is not a dictionary word", start)
		return nil, err
	}

	if !f.dict.Contains(end) {
		err := fmt.Errorf("ending word %s is not a dictionary word", end)
		return nil, err
	}

	if start == end {
		return []string{start}, nil
	}

	left := Route{start: []string{}}
	right := Route{end: []string{}}

	result, err := f._findPath(left, right)
	if err != nil {
		fmt.Println("do something")
	}

	return result, nil
}

// Find the shortest path between two routes.
func (f Finder) _findPath(left Route, right Route) ([]string, error) {
	if len(left) == 0 || len(right) == 0 {
		return nil, fmt.Errorf("no match")
	}

	match := checkMatch(left, right)

	if match != "" {
		return makePath(left, right, match)
	}

	newRoute := make(Route)

	if len(left) <= len(right) {
		for k, v := range left {
			mutants := f.dict.Mutate(k)
			for mutant := range mutants {
				if !left.PathContains(k, mutant) {
					newRoute[mutant] = append(v, k)
				}
			}
		}
		return f._findPath(newRoute, right)
	} else {
		for k, v := range right {
			mutants := f.dict.Mutate(k)
			for mutant, _ := range mutants {
				if !right.PathContains(k, mutant) {
					newRoute[mutant] = append(v, k)
				}
			}
		}
		return f._findPath(left, newRoute)
	}
}

// Determine if there are any shared words between two routes.
func checkMatch(left Route, right Route) string {
	for k, _ := range left {
		if _, ok := right[k]; ok {
			return k
		}
	}
	return ""
}

// Glue two routes (string arrays) together with an intervening word.
// The second array is reversed prior to concatenation.
// For example:
// left: [one, two], right: [shoe, my], word: buckle => [one, two, buckle, my, shoe]
func makePath(left Route, right Route, word string) ([]string, error) {
	newPath := append(left[word], word)
	newPath = append(newPath, reverseArray(right[word])...)
	return newPath, nil
}

// Reverse an array of strings.
func reverseArray(words []string) []string {
	wordsLen := len(words)
	reversed := make([]string, wordsLen)

	for i := 0; i < wordsLen; i++ {
		reversed[wordsLen-i-1] = words[i]
	}

	return reversed
}
