/**
Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.
Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000
**/

class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
		
        for(int i =0;i<nums.length;i++){
			if(i ==0)
				continue;
			else if(nums[i]!=0 && nums[i-1]!=0)
				nums[i]=nums[i]+nums[i-1];
			else 
				continue;
		}
		int max=nums[0];
		for(int j =1;j<nums.length;j++){
			if(nums[j]>max)
				max = nums[j];
		}
		return max;
    }
}