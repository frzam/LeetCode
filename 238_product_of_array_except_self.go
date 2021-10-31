func productExceptSelf(nums []int) []int {
    count := 0
    total := 1
    
    for i:=0;i<len(nums);i++{
        if nums[i] != 0{
            total *= nums[i]
        }else {
            count ++
        }
    }
    if count > 1 {
        return make([]int, len(nums))
    }
    for i := 0;i<len(nums);i++{
        if nums[i] == 0 {
            nums[i] = total   
        }else if count == 0{
            nums[i] = total/nums[i]
        }else{
            nums[i] = 0
        }
    }
    return nums
}
