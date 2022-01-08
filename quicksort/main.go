package main

import "fmt"

func main() {
	numbers := []int{3, 5, 9, 133, 44, 20, 1, 7}

	fmt.Printf("Result of the sum: %v\n", sumOf(numbers))
	fmt.Printf("Result of the count: %v\n", countTotalItemsIn(numbers))
	fmt.Printf("Result of the highest number: %v\n", highestNumberFrom(numbers))
	fmt.Printf("Result of the sorted list: %v\n", sortOf(numbers))
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

func highestNumberFrom(numbers []int) int {
	if len(numbers) == 2 {
		if numbers[0] > numbers[1] {
			return numbers[0]
		} else {
			return numbers[1]
		}
	} else {
		subMax := highestNumberFrom(numbers[1:])

		if numbers[0] > subMax {
			return numbers[0]
		} else {
			return subMax
		}
	}
}

func sortOf(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	middle := []int{numbers[len(numbers)/2]}

	firstHalf := numbers[:(len(numbers) / 2)]
	secondHalf := numbers[(len(numbers)/2)+1:]

	var lowers []int
	var highers []int

	for _, elem := range firstHalf {
		if elem > middle[0] {
			highers = append(highers, elem)
		} else {
			lowers = append(lowers, elem)
		}
	}

	for _, elem := range secondHalf {
		if elem > middle[0] {
			highers = append(highers, elem)
		} else {
			lowers = append(lowers, elem)
		}
	}

	return append(sortOf(lowers), append(middle, sortOf(highers)...)...)
}
