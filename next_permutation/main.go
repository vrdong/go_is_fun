package main

import (
	"sort"
)

func main() {
	nextPermutation([]int{3, 2, 1})
}

func nextPermutation(nums []int) {
	isChange := false
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			min := i + 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > nums[i] && nums[j] < nums[min] {
					min = j
				}
			}
			swap(nums, i, min)
			sort.Ints(nums[i+1:])
			isChange = true
			break
		}
	}
	if !isChange {
		sort.Ints(nums)
	}
	//fmt.Println(nums)
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
