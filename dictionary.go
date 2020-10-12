package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type DictionaryService interface {
	Contains(word string) bool
	Mutate(word string) Dictionary
	Add(word string)
}

// Using a map here due to having the key lookup complexity of O(1), instead of O(n) lookup complexity
// in an array.
type Dictionary map[string]bool

// Create a new file based dictionary.  The dictionary should be a plain text file with one
// word per line.
func NewFileDictionary(fileName string) Dictionary {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("unable to open dictionary file %v", fileName)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	wordMap := make(Dictionary)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) > 0 {
			wordMap[word] = true
		}
	}

	return wordMap
}

// Adds a word to the dictionary.
func (d Dictionary) Add(word string) {
	d[word] = true
}

// Does the word exist in the dictionary?
func (d Dictionary) Contains(word string) bool {
	_, ok := d[word]

	return ok
}

// Computes all valid mutations of a word.  Valid mutations are all only a single letter change,
// either added, deleted, or changed within the word.  A word must be in the dictionary
// to be considered valid, and a word cannot mutate into itself.
func (d Dictionary) Mutate(word string) Dictionary {
	letters := _getLetters()
	mutations := make(Dictionary)
	var mutant string

	for i := range word {
		mutant = word[:i] + word[i+1:]
		if d.Contains(mutant) {
			mutations[mutant] = true
		}

		for _, c := range letters {
			mutant = word[:i] + c + word[i:]
			if d.Contains(mutant) {
				mutations[mutant] = true
			}

			mutant = word[:i] + c + word[i+1:]
			if d.Contains(mutant) {
				mutations[mutant] = true
			}
		}
	}

	for _, c := range letters {
		mutant = word + c
		if d.Contains(mutant) {
			mutations[mutant] = true
		}
	}

	delete(mutations, word)

	return mutations
}

// Get an array of lowercase letters from 'a' to 'z'
func _getLetters() []string {
	var letters []string

	for i := 'a'; i <= 'z'; i++ {
		letters = append(letters, fmt.Sprintf("%c", i))
	}
	return letters
}
