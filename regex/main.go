package main

import "fmt"

func main() {
	isMatch("a", ".*..a*")
}

func isMatch(s string, p string) bool {
	m := len(p)
	n := len(s)
	A := make([][]bool, m+1, m+1)
	for i := range A {
		A[i] = make([]bool, n+1, n+1)
	}

	A[0][0] = true
	for i := 1; i <= m; i++ {
		switch p[i-1] {
		case '*':
			A[i][0] = A[i-2][0]
		case '.':
			A[i][0] = false
		default:
			A[i][0] = false
		}
	}

	for j := 1; j <= n; j++ {
		A[0][j] = false
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			switch p[i-1] {
			case '*':
				conditionMatch := p[i-2] == s[j-1] || p[i-2] == '.'
				A[i][j] = (A[i-2][j-1] && conditionMatch) || (A[i-2][j]) || (A[i][j-1] && conditionMatch)
			case '.':
				A[i][j] = A[i-1][j-1]
			default:
				A[i][j] = p[i-1] == s[j-1] && A[i-1][j-1]
			}
		}
	}
	for i := -1; i <= n; i++ {
		if i < 1 {
			fmt.Printf("0     ")
		} else {
			fmt.Printf("%c     ", s[i-1])
		}
	}
	fmt.Println()
	for i := 0; i <= m; i++ {
		for j := -1; j <= n; j++ {
			if j == -1 {
				if i == 0 {
					fmt.Printf("0     ")
				} else {
					fmt.Printf("%c     ", p[i-1])
				}
			} else {
				if A[i][j] == false {
					fmt.Printf("false ")
				} else {
					fmt.Printf("true  ")
				}

			}
		}
		fmt.Println()
	}
	return A[m][n]
}
