package main

import "fmt"

func main() {
	fmt.Print(intToRoman(3999))
}

func intToRoman(num int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	div := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	divIndex := 0
	for num != 0 {
		if num >= div[divIndex] {
			num -= div[divIndex]
			result = result + roman[divIndex]
		} else {
			divIndex += 1
		}
	}
	return result
}
