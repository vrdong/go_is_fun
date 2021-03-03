package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{6, 6, 6, 6, 6}))
}

func distributeCandies(candyType []int) int {
	m := make(map[int]bool)
	numberOfType := 0
	for i := 0; i < len(candyType); i++ {
		if m[candyType[i]] == false {
			m[candyType[i]] = true
			numberOfType++
		}
	}
	if numberOfType < (len(candyType) / 2) {
		return numberOfType
	} else {
		return len(candyType) / 2
	}
}
