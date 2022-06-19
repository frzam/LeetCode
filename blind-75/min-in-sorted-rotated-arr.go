func findMin(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	m := 0
	l := 0
	h := len(nums) - 1
	for l < h {
		m = (l + h) / 2
		if nums[m] > nums[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
