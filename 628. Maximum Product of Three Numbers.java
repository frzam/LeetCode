
/**
Given an integer array, find three numbers whose product is maximum and output the maximum product.

Example 1:

Input: [1,2,3]
Output: 6
 

Example 2:

Input: [1,2,3,4]
Output: 24
 

Note:

The length of the given array will be in range [3,104] and all elements are in the range [-1000, 1000].
Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.

Approach:
Sort the nums[] using count sort, since -1000<nums[i]<1000.
Check the multiplication of first two i =0,1 and second last two i.e i=n-2, n-3.
return last multiplied by which among the above two is higher.

Solution Analysis:
Runtime: 2 ms, faster than 99.69% of Java online submissions for Maximum Product of Three Numbers.
Memory Usage: 38.8 MB, less than 100.00% of Java online submissions for Maximum Product of Three Numbers.

**/
class Solution {
    public int maximumProduct(int[] nums) {
      int countArr[] = new int[2001];
        for(int i =0;i<nums.length;i++){
            countArr[nums[i]+1000]++;
        }
        int k=0;
        int i =0;
        while(i<countArr.length){
            if(countArr[i]>0){
                countArr[i]--;
                nums[k++]=i-1000;
            }
            if(countArr[i]==0)
                i++;
        }
        if(nums[0]*nums[1]>nums[nums.length-3]*nums[nums.length-2])
            return nums[0]*nums[1]*nums[nums.length-1];
        else 
            return nums[nums.length-3]*nums[nums.length-2]*nums[nums.length-1];
    }
}