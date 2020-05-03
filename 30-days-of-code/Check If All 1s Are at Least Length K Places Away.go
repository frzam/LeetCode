func kLengthApart(nums []int, k int) bool {
    seen := false
    count := 0
    for i := 0;i<len(nums);i++{
        if !seen && nums[i] == 1{
            seen = true
        }else if seen && nums[i] != 1{
            count ++
        }else if seen && nums[i] == 1{
            if count < k{
                return false
            }
            count = 0
        }
    }
    return true
}
