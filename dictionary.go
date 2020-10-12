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
}

type Dictionary map[string]bool

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

func (d Dictionary) Add(word string) {
	d[word] = true
}

func (d Dictionary) Contains(word string) bool {
	_, ok := d[word]

	return ok
}

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

func _getLetters() []string {
	var letters []string

	for i := 'a'; i <= 'z'; i++ {
		letters = append(letters, fmt.Sprintf("%c", i))
	}
	return letters
}
