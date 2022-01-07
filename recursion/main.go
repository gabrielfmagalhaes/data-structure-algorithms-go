package main

import "fmt"

func main() {
	num := 5

	result := factorialOf(num)

	fmt.Printf("Factorial of %v is %v\n", num, result)
}

func factorialOf(num int) int {
	if num == 1 {
		return num
	}

	return num * factorialOf(num-1)
}