package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		threeResult := threeSum(nums[i+1:], target-nums[i])
		for j := 0; j < len(threeResult); j++ {
			miniResult := []int{nums[i]}
			result = append(result, append(miniResult, threeResult[j]...))
		}
	}
	return result
}

func threeSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	//sort.Ints(nums)
	for index, value := range nums {
		if index != 0 && value == nums[index-1] {
			continue
		}
		left := index + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + value
			switch {
			case sum < target:
				left += 1
			case sum > target:
				right -= 1
			default:
				result = append(result, []int{value, nums[left], nums[right]})
				left += 1
				for nums[left] == nums[left-1] && left < right {
					left += 1
				}
			}
		}
	}
	return result
}
