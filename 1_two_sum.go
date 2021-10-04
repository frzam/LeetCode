package leetcode

func twoSum(nums []int, target int) []int {

	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := count[target-nums[i]]; ok {
			return []int{i, count[target-nums[i]]}
		}
		count[nums[i]] = i

	}
	return []int{}
}
