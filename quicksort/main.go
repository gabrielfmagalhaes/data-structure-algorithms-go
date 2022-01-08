package main

import "fmt"

func main() {
	numbers := []int{3, 5, 9, 133, 44, 20, 1, 7}

	fmt.Printf("Result of the sum: %v\n", sumOf(numbers))
}

func sumOf(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	} else {
		return numbers[0] + sumOf(numbers[1:])
	}
}
