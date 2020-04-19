func search(nums []int, target int) int {
    return binarySearch(nums,0,len(nums)-1,target)
}

func binarySearch(nums []int, l , r,target int)int{
    if l <=r {
        mid := (l+r)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid] > target{
            return binarySearch(nums,l,mid-1,target)
        }else{
            return binarySearch(nums,mid+1,r,target)
        } 
    }
    return -1
}
