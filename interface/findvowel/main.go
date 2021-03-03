package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (s MyString) FindVowels() []rune {
	var result []rune
	for _, character := range s {
		if character == 'a' || character == 'i' || character == 'e' || character == 'u' || character == 'o' {
			result = append(result, character)
		}
	}
	return result
}

func main() {
	var v VowelsFinder
	var s MyString
	s = "Cao Ba Dong"
	v = s
	fmt.Printf("%c", v.FindVowels())
}
