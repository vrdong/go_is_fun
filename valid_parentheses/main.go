package main

import "fmt"

func main() {
	fmt.Println(isValid(""))
}

func isValid(s string) bool {
	var stack string
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = stack + "("
		case '{':
			stack = stack + "{"
		case '[':
			stack = stack + "["
		case ')':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}
	return true
}
