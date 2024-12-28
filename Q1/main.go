package main

import (
	"fmt"
	"sort"
	"strings"
)

func countA(word string) int {
	return strings.Count(word, "a") + strings.Count(word, "A")

}

func compareWords(word1, word2 string) bool {
	count1 := countA(word1)
	count2 := countA(word2)

	if count1 > count2 {
		return len(word1) > len(word2)
	}
	return count1 > count2
}

func orderWords(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return compareWords(words[i], words[j])
	})
	return words

}

func main() {
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	orderedWords := orderWords(words)
	fmt.Println(orderedWords)
}
