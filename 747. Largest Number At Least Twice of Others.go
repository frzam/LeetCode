/**
In a given integer array nums, there is always exactly one largest element.

Find whether the largest element in the array is at least twice as much as every other number in the array.

If it is, return the index of the largest element, otherwise return -1.

Example 1:

Input: nums = [3, 6, 1, 0]
Output: 1
Explanation: 6 is the largest integer, and for every other number in the array x,
6 is more than twice as big as x.  The index of value 6 is 1, so we return 1.
 

Example 2:

Input: nums = [1, 2, 3, 4]
Output: -1
Explanation: 4 isn't at least as big as twice the value of 3, so we return -1.
 

Note:

nums will have a length in the range [1, 50].
Every nums[i] will be an integer in the range [0, 99].

Analysis:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Largest Number At Least Twice of Others.
Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Largest Number At Least Twice of Others.
*/

func dominantIndex(nums []int) int {
    max,secMax := math.MinInt32,math.MinInt32
    loc := -1
    for i:=0;i<len(nums);i++{
        if nums[i] >= max{
            secMax = max
            max = nums[i]
            loc = i
        }else if nums[i] >secMax{
            secMax = nums[i]
       }
    }
    if max >= secMax*2{
        return loc
    }else{
        return -1
    }
    return loc
}