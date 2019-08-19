/**
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
**/

class Solution {
    public List<Integer> findDisappearedNumbers(int[] nums) {
        ArrayList<Integer> list = new ArrayList<Integer>();
		int count[] = new int[nums.length];
		for(int i=0;i<nums.length;i++){
			count[nums[i]-1]++;
		}
		for(int j=0;j<count.length;j++){
			if(count[j] == 0)
				list.add(j+1);
		}
		return list;
    }
}