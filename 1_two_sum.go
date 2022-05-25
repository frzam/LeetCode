package leetcode

func twoSum(nums []int, target int) []int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := count[nums[i]]; ok {
			return []int{v, i}
		} else {
			count[target-nums[i]] = i
		}
	}
	return []int{-1, -1}
}
