package leetcode

import "math"

func maxSubArray(arr []int) int {
	max := math.MinInt64
	sum := 0
	for i := 0; i < len(arr); i++ {
		if sum < 0 {
			sum = arr[i]
		} else {
			sum += arr[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
