package main

import (
	"fmt"
)

func main() {
	maxNum := 1024
	expect := 22

	sortList := generateRandomList(maxNum)

	findNumber(expect, sortList)
}

func generateRandomList(num int) []int {
	list := make([]int, num)

	for i := 0; i < num; i++ {
		list[i] = i
	}

	return list
}

func findNumber(guessNumber int, list []int) {
	low := 0
	high := list[len(list)-1]

	count := 1

	for low <= high {

		middle := (low + high) / 2
		guessed := list[middle]

		if guessed == guessNumber {
			fmt.Println("Number guessed: ", guessed)
			return
		}
		if guessed > guessNumber {
			high = middle - 1
		} else {
			low = middle + 1
		}

		fmt.Println("Attempt: ", count)
		count++
	}
}
