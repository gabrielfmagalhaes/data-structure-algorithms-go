package main

import "fmt"

func main() {
	numbers := []int{3, 5, 9, 133, 44, 20, 1, 7}

	fmt.Printf("Result of the sum: %v\n", sumOf(numbers))
	fmt.Printf("Result of the count: %v\n", countTotalItemsIn(numbers))
}

func sumOf(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	} else {
		return numbers[0] + sumOf(numbers[1:])
	}
}

func countTotalItemsIn(numbers []int) int {
	if len(numbers) == 1 {
		return 1
	} else {
		return 1 + countTotalItemsIn(numbers[1:])
	}
}
