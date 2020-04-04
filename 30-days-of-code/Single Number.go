func singleNumber(nums []int) int {
    countMap := make(map[int]int)
    for i:=0;i<len(nums);i++{
        countMap[nums[i]]++
    }
    for k,v := range countMap{
        if v == 1{
            return k
        }
    }
    return -1
}