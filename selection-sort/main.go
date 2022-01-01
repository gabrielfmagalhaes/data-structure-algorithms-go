package main

import "fmt"

func main() {
	unsortedList := initUnsortedList()
	sortedList := sortList(unsortedList)

	printList(sortedList)
}

func initUnsortedList() map[string]int {
	return map[string]int{"Gorillaz": 141,
		"Bo Burnham":     234,
		"Lil Nas X":      94,
		"Kendrick Lamar": 88,
		"David Bowie":    61,
		"Elton John":     111}
}

func sortList(unsortedList map[string]int) map[string]int {
	sortedList := make(map[string]int)
	listLength := len(unsortedList)

	for i := 0; i < listLength; i++ {
		higherKey, higherValue := findHigher(unsortedList)

		sortedList[higherKey] = higherValue

		delete(unsortedList, higherKey)
	}

	return sortedList
}

func findHigher(unsortedList map[string]int) (resultKey string, resultValue int) {
	resultKey = ""
	resultValue = -1

	for key, value := range unsortedList {
		if resultValue < value {
			resultValue = value
			resultKey = key
		}
	}

	return resultKey, resultValue
}

func printList(sortedList map[string]int) {
	for key, value := range sortedList {
		fmt.Printf("Artist: %s, amount of plays: %v\n", key, value)
	}
}
