package main

import "fmt"

func main() {
	recursiveFunction(1)
}

func recursiveFunction(current int) {
	if current > 3 {
		return
	}
	if current == 1 {
		fmt.Println(current + 1)
	} else {
		fmt.Println(current * current)
	}
	recursiveFunction(current + 1)
}
