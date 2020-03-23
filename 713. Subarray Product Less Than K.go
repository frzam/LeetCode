func numSubarrayProductLessThanK(nums []int, k int) int {
    p := 1
    count := 0
    i:=0
    for j:=0;j<len(nums);j++{
        p *= nums[j]
        for i<= j&& p>=k{
            p /= nums[i]
            i++
        }
        count += j - i + 1
    }
    return count
}