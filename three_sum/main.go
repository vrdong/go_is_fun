package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Print(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Print(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for index, value := range nums {
		if index != 0 && value == nums[index-1] {
			continue
		}
		left := index + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + value
			switch {
			case sum < 0:
				left += 1
			case sum > 0:
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

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := 9999999
	result := 0
	for index, value := range nums {
		left := index + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + value
			switch {
			case sum < target:
				if min > target-sum {
					min = target - sum
					result = sum
				}
				left += 1
			case sum > target:
				if min > sum-target {
					min = sum - target
					result = sum
				}
				right -= 1
			default:
				return target
			}
		}
	}
	return result
}
