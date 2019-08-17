/**
iven a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

Example 1:
Input: [1, 2, 2, 3, 1]
Output: 2
Explanation: 
The input array has a degree of 2 because both elements 1 and 2 appear twice.
Of the subarrays that have the same degree:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
The shortest length is 2. So return 2.
Example 2:
Input: [1,2,2,3,1,4,2]
Output: 6
Note:

nums.length will be between 1 and 50,000.
nums[i] will be an integer between 0 and 49,999.
Solution Analysis:
Runtime: 9 ms, faster than 94.98% of Java online submissions for Degree of an Array.
Memory Usage: 40.5 MB, less than 94.44% of Java online submissions for Degree of an Array.
**/

class Solution {
    public int findShortestSubArray(int[] nums) {
        HashMap<Integer, int[]> degreeMap = new HashMap<Integer,int[]>();
        for(int i =0;i<nums.length;i++){
            if(!degreeMap.containsKey(nums[i])){
                degreeMap.put(nums[i], new int[]{1,i,i});
            }
            else{
                int[] temp = degreeMap.get(nums[i]);
                temp[0]++;
                temp[2] = i;
            }
        }
        int degree = Integer.MIN_VALUE, len = Integer.MAX_VALUE;
        for(int[] value : degreeMap.values()){
            if(degree<value[0])
            {
                degree= value[0];
                len = value[2]-value[1]+1;
            }
            else if(degree == value[0]){
                len = Math.min((value[2]-value[1]+1), len);
            }
        }
      return len;  
    }
}