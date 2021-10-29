func longestConsecutive(nums []int) int {
    count := make(map[int]int)
    max := 0
    for i:=0;i<len(nums);i++{
        if _, ok := count[nums[i]]; !ok {
            left, _ := count[nums[i]-1]
            right, _ := count[nums[i]+ 1]
            sum := left + right + 1
            count[nums[i]]= sum
            count[nums[i] - left] = sum
            count[nums[i] + right] = sum
            if sum > max {
                max = sum
            } 
        }

    }
    return max
}
