package main

import "fmt"

func main() {
	nums := []int{}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	current := 0
	for i := 1; i < len(nums); i++ {
		if nums[current] == val {
			if nums[i] != val {
				tmp := nums[current]
				nums[current] = nums[i]
				nums[i] = tmp
				current++
			}
		} else {
			current++
		}
	}
	if nums[current] == val {
		return current
	}

	return current + 1
}
