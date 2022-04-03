package main

import "fmt"

var brackets = map[string]string{
	")": "(",
	"]": "[",
	"}": "{"}

func main() {
	s := "({[]{}}())[]"

	result := isParenthesesExpressionValid(s)

	fmt.Printf("Expression %s is %v", s, result)
}

func isParenthesesExpressionValid(s string) bool {
	var stack []string

	for i := 0; i < len(s); i++ {
		char := string(s[i])

		if char == "(" || char == "[" || char == "{" {
			stack = append(stack, char)
		} else {
			if len(stack) <= 0 {
				return false
			}

			n := len(stack) - 1

			bracket := getBracket(char)

			if bracket == "" {
				return false
			}

			if bracket != stack[n] {
				return false
			}

			stack = stack[:n]
		}
	}

	return len(stack) <= 0
}

func getBracket(char string) string {
	return brackets[char]
}
