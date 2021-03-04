package main

import "fmt"

/*a array nums contains n distinct numbers in range [0,n]*/
func main() {
	fmt.Println(missingNumber([]int{9,6,4,2,3,5,7,0,1}))
}


func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return len(nums) * (len(nums) + 1) / 2 - sum

}
