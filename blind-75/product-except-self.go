package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	r := 1
	for j := len(nums) - 2; j >= 0; j-- {
		r = r * nums[j+1]
		res[j] = res[j] * r
	}

	return res
}
