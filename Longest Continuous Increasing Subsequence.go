func findLengthOfLCIS(nums []int) int {
    if len(nums)==0{
        return 0
    }
    res := 1
    count := 1
    for i:=0;i<len(nums)-1;i++{
        if nums[i] < nums[i+1]{
             count ++
             if  count > res{
                 res = count
            }
        }else {
             count = 1
        }
    }
    return res
}
