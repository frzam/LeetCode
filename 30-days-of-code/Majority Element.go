func majorityElement(nums []int) int {
    n := nums[0]
    count := 1
    for i:=1;i<len(nums);i++{
        if count == 0{
            n = nums[i]
        }
        if n == nums[i]{
            count ++
        }else{
            count --
        }
    }
    return n
}
