/**
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].



**/

class Solution {
    public int[] twoSum(int[] nums, int target) {
        int ret[] = new int[2];
        int arr[] = new int[nums.length];
        
        int a = 0;
        int b = 0;
        int i = 0;
        int j = nums.length-1;
        for (int k=0; k<nums.length; k++) 
            arr[k] = nums[k];
        Arrays.sort(nums);
        while(true){
            if(nums[i] + nums[j] > target )
                j--;
            else if (nums[i] + nums[j] < target)
                i++;
            else if (nums[i]+ nums[j] == target)
            {
                a = nums[i];
                b = nums[j];
               
                break;
            }
       }
        boolean flag = false;
        for(int l = 0; l <nums.length;l++){
            if (a == arr[l] && flag == false){
                ret[0] = l;
                flag  = true;
                continue;
            }
            if (b == arr[l])
                ret[1] = l;
        }
        return ret;
    }
}