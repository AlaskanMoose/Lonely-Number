package main

import "testing"

func TestFindLonelyNumber(t *testing.T) {
	var a = [9]int{2,3,4,1,5,3,5,2,4}

	actualResult := findLonelyNumber(a[:]...)
	var expectedResult = 1

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}
}