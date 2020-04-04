package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sum := nums[0]
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if sum > max {
			max = sum
		}
		sum += nums[i]
		if sum < 0 {
			sum = 0
		}

	}
	return max
}

func main() {
	nums := []int{-2, -1}
	num := maxSubArray(nums)
	fmt.Println(num)
}
