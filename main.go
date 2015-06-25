package main

import (
	"fmt"
	"github.com/YurkoHoshko/coursera/week1"
)

func main() {
	unsortedArray := []int{2, 4, 6, 1, 3, 5}
	sortedArray, inversions := week1.FindInversions(unsortedArray)
	var inversionsCount int

	for _, v := range inversions {
		inversionsCount += len(v)
	}

	fmt.Println("Input array: ", unsortedArray)
	fmt.Println("Sorted array: ", sortedArray)
	fmt.Println("Count of inversions:", inversionsCount)
}
