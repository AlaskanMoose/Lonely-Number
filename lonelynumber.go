package main

import "fmt"

func main() {
	fmt.Printf("lonely number\n")
}
func findLonelyNumber(a ...int) int {

	set := make(map[int]bool)

	for _, num := range a {
		set[num] = true
	}
	var doubleTheAmount int

	for key := range set {
		doubleTheAmount += key * 2
	}

	for _, num := range a {
		doubleTheAmount -= num
	}
	return doubleTheAmount
}