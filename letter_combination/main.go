package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

var digitLetter = map[byte][]string{
	'0': {},
	'1': {},
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	switch len(digits) {
	case 1:
		return digitLetter[digits[0]]
	case 0:
		return []string{}
	default:
		var result []string
		for i := 0; i < len(digitLetter[digits[0]]); i++ {
			subResult := letterCombinations(digits[1:])
			for j := 0; j < len(subResult); j++ {
				result = append(result, digitLetter[digits[0]][i]+subResult[j])
			}
		}
		return result
	}
}

//["ad", "ae", "af", "bd", "be","bf", "cd", "ce", "cf"]
