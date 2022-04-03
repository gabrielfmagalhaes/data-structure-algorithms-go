package main

import (
	"fmt"
)

var romans = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000}

func main() {
	input := "MCMXCVII"

	resultConverted := romanToInteger(input)

	fmt.Printf("Converted roman value %s to %d", input, resultConverted)
}

func romanToInteger(romanExpression string) int {
	prev := ""
	result := 0

	for i := 0; i < len(romanExpression); i++ {
		for char, value := range romans {
			if char == string(romanExpression[i]) {
				if prev == "I" && (char == "V" || char == "X") {
					result += value - 2

					break
				} else if prev == "X" && (char == "L" || char == "C") {
					result += value - 20

					break
				} else if prev == "C" && (char == "D" || char == "M") {
					result += value - 200

					break
				} else {
					result += value

					break
				}
			}
		}

		prev = string(romanExpression[i])
	}

	return result
}
