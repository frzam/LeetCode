func findDuplicates(nums []int) []int {
   var res []int
    for i:= 0;i<len(nums);i++{
        idx := abs(nums[i])-1
        if nums[idx]<0{
            res = append(res, idx + 1)
        }
         nums[idx] *= -1
        
    }
    return res
}
func abs(n int)int{
    if n <0{
        return n * -1
    }
    return n
}
