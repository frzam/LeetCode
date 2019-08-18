/**
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

Input: [3,0,1]
Output: 2
Example 2:

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?
Solution Analysis:
Runtime: 0 ms, faster than 100.00% of Java online submissions for Missing Number.
Memory Usage: 39 MB, less than 100.00% of Java online submissions for Missing Number.

**/
class Solution {
    public int missingNumber(int[] nums) {
      int sum =0;
        for(int i =0;i<nums.length;i++)
            sum +=nums[i];
        return (nums.length)*(nums.length+1)/2 - sum;
    }
}



class Solution {
    public int missingNumber(int[] nums) {
       int countFreq[] = new int[nums.length+1];
        for(int i =0;i<nums.length;i++)
            countFreq[nums[i]]++;
        for(int j =0;j<countFreq.length;j++){
            if(countFreq[j] ==0)
                return j;
        }
        return -1;
    }
}