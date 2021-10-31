func search(nums []int, target int) int {
    return bs(nums, target, 0, len(nums)-1)
}

func bs(nums []int, target, l, h int) int{
    if l <= h {
        m := (l + h)/2
        if nums[m] == target{
            return m
        }
        if nums[l] <= nums[m]{
            if nums[m] > target && nums[l] <= target{
                return bs(nums, target, l, m -1)
            }else {
                return bs(nums, target, m+1, h)
            }
        }else{
            if nums[m] < target && nums[h] >= target {
                return bs(nums, target, m+1, h)
            }else {
                return bs(nums, target, l, m -1)
            }
        }

    }
    return -1
}
