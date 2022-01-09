package main

import "fmt"

func main() {
	unsortedList := []int{141, 234, 94, 88, 61, 111}

	fmt.Printf("Sorted list: %v\n", sortListFrom(unsortedList))
}

func sortListFrom(unsortedList []int) []int {
	var sortedList []int

	for _, _ = range unsortedList {
		highestIndex := highestNumberFrom(unsortedList)

		sortedList = append(sortedList, unsortedList[highestIndex])

		unsortedList = append(unsortedList[:highestIndex], unsortedList[highestIndex+1:]...)
	}

	return sortedList
}

func highestNumberFrom(numbers []int) int {
	highestNumber := numbers[0]
	highestIndex := 0

	for index, elem := range numbers {
		if highestNumber < elem {
			highestNumber = elem
			highestIndex = index
		}
	}

	return highestIndex
}
