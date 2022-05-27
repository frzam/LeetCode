package blind75

func containsDuplicate(nums []int) bool {
	count := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := count[nums[i]]; ok {
			return true
		}
		count[nums[i]] = struct{}{}
	}
	return false
}
