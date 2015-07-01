package main

import (
	"bufio"
	"fmt"
	"github.com/YurkoHoshko/coursera/week1"
	"os"
	"strconv"
)

func main() {
	runWeek1()
}

func runWeek1() {
	unsortedArray := []int{}
	file, err := os.Open("resources/ProgrammingTask1.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		unsortedArray = append(unsortedArray, number)
	}

	sortedArray, inversionsCount := week1.FindInversions(unsortedArray)
	fmt.Println("Input array: ", unsortedArray)
	fmt.Println("Sorted array: ", sortedArray)
	fmt.Println("Count of inversions:", inversionsCount)
}
