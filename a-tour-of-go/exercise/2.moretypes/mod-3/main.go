package main

import "fmt"

func countWord(words []string) map[string]int {
	m := make(map[string]int)
	for _, v := range words {
		m[v]++
	}
	return m
}

func mostCommon(counts map[string]int) (word string, count int) {

	var maxString string
	var maxValue int

	for i, v := range counts {
		if v > maxValue {
			maxString = i
			maxValue = v
		}
	}

	return maxString, maxValue

}

func removeWord(counts map[string]int, word string) bool {

	_, ok := counts[word]

	if !ok {
		return false
	}

	delete(counts, word)

	return true

}

func main() {
	words := []string{"go", "code", "go", "learn", "code", "go", "map", "learn"}

	counts := countWord(words)

	fmt.Println(words)
	fmt.Println(counts)
	fmt.Println(mostCommon(counts))
	fmt.Println(removeWord(counts, "code"))
	fmt.Println(counts)

}
