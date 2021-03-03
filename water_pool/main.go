package main

import "fmt"

func main() {
	fmt.Print(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	m := len(height)
	Max := 0
	maxStart := 0
	for i := 0; i < m-1; i++ {
		fmt.Println(Max)
		if height[maxStart] >= height[i] && i != 0 {
			continue
		} else {
			maxEndI := 0
			currentMaxI := 0
			for j := i + 1; j < m; j++ {
				if height[j] >= maxEndI {
					maxEndI = height[j]
					currentMaxI = (j - i) * min(height[i], height[j])
				} else {
					if currentMaxI < (j-i)*min(height[i], height[j]) {
						maxEndI = height[j]
						currentMaxI = (j - i) * min(height[i], height[j])
					}
				}
			}
			if currentMaxI > Max {
				Max = currentMaxI
				maxStart = i
			}
		}
	}
	return Max
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
