package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	box := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		box = max(box, 0) + nums[i]
		ans = max(ans, box)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	input := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(input))
}
