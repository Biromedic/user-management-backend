package main

import "fmt"

func main() {
	data := []string{"apple", "pie", "apple", "red",
		"red", "red", "a", "a", "a", "BANANA", "BANANA", "BANANA", "BANANA", "BANANA",
		"hello", "world", "hello", "go", "go", "go", "go", "go",
		"x", "y", "z", "z", "y", "y"}

	result := getMostRepeatedItem(data)

	fmt.Printf("\"%s\"\n", result)
}

func getMostRepeatedItem(data []string) string {
	countMap := countOccurrences(data)

	for item, count := range countMap {
		fmt.Printf("%s: %d\n", item, count)
	}

	return findMostRepeated(countMap)
}

func findMostRepeated(countMap map[string]int) string {
	var mostRepeated string
	maxCount := 0
	for key, value := range countMap {
		if value > maxCount {
			maxCount = value
			mostRepeated = key
		}
	}
	return mostRepeated
}

func countOccurrences(data []string) map[string]int {
	countMap := make(map[string]int)
	for _, item := range data {
		countMap[item]++
	}
	return countMap

}
