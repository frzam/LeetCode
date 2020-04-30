func createTargetArray(nums []int, index []int) []int {
    var res []int
    for i := range nums{
        res = append(res[:index[i]], append([]int{nums[i]}, res[index[i]:]...)...)
     }
    return res
    
}
