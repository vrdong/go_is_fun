package main

import "fmt"

func main() {
	fmt.Println(divide(-9, -5))
}

func divide(dividend int, divisor int) int {
	isNegativeEnd := dividend < 0
	isNegativeSor := divisor < 0
	isDiffSign := isNegativeEnd != isNegativeSor

	count := 0
	if isDiffSign {
		if isNegativeEnd {
			for dividend+divisor <= 0 {
				count -= 1
				dividend = dividend + divisor
			}
		} else {
			for dividend+divisor >= 0 {
				count -= 1
				dividend = dividend + divisor
			}
		}
	} else {
		if isNegativeEnd {
			for dividend-divisor <= 0 {
				count += 1
				dividend = dividend - divisor
			}
		} else {
			for dividend-divisor >= 0 {
				count += 1
				dividend = dividend - divisor
			}
		}
	}
	return count
}
