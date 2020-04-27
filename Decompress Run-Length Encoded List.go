func decompressRLElist(nums []int) []int {
    var res []int
    for i:=0;i<len(nums)-1;i=i+2{
        j := i+1
        for k:=0;k<nums[i];k++{
            res = append(res, nums[j])
        }
    }
    return res
}
