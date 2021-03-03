package main

import "fmt"

func main() {
	nums := []int{}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	current := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			current += 1
			nums[current] = nums[i]
		}
	}
	nums = nums[:current+1]
	return current + 1
}
