func searchRange(nums []int, target int) []int {
    return []int{firstBS(nums, target, 0, len(nums)-1), secondBS(nums, target, 0, len(nums)-1)}
}

func firstBS(nums []int, target, low, high int) int{
    idx := -1
    for low <= high {
        mid := (high+ low)/2
        
        if nums[mid] == target {
            idx = mid
        }
        if nums[mid] >= target {
            high = mid - 1
        }else {
            low =  mid + 1
        }
    }   
    return idx
}

func secondBS(nums[]int, target, low, high int) int {
    idx := -1
    for low <= high{
        mid := (low+high)/2
        if nums[mid] == target{
            idx = mid
        }
        if nums[mid]<=target{
            low = mid + 1
        }else{
            high = mid - 1
        }
    }
    return idx
}
