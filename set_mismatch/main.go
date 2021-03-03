package main

import "fmt"

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}

func findErrorNums(nums []int) []int {
	x := make([]bool, len(nums), len(nums))
	sum := 0
	dup := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if x[nums[i]-1] == false {
			x[nums[i]-1] = true
		} else {
			dup = nums[i]
		}
	}
	diff := sum - len(nums)*(len(nums)+1)/2
	result := make([]int, 2)
	result[0] = dup
	result[1] = dup - diff
	return result
}
