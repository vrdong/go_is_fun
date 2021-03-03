package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{}))
}

func longestCommonPrefix(strs []string) string {
	min := 999
	minValue := ""
	for _, value := range strs {
		if min > len(value) {
			min = len(value)
			minValue = value
		}
	}

	for i := len(minValue); i >= 0; i-- {
		if isPrefix(minValue[:i], strs) {
			return minValue[:i]
		}
	}
	return ""
}

func isPrefix(pre string, strs []string) bool {
	for i := range strs {
		if pre != strs[i][:len(pre)] {
			return false
		}
	}
	return true
}
