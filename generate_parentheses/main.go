package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(10))
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	var doer func(int, int, int, string)

	doer = func(n int, currentOpen int, numOpen int, current string) {
		if len(current) == n*2 {
			result = append(result, current)
		}

		for i := 0; i <= 1; i++ {
			if i == 0 {
				if numOpen < n {
					doer(n, currentOpen+1, numOpen+1, current+"(")
				}
			} else {
				if currentOpen > 0 {
					doer(n, currentOpen-1, numOpen, current+")")
				}
			}
		}
	}
	doer(n, 0, 0, "")
	return result
}
