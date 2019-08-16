/**
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Note:

Each element in the result must be unique.
The result can be in any order.
 
Solution Analysis:
Runtime: 2 ms, faster than 98.40% of Java online submissions for Intersection of Two Arrays.
Memory Usage: 37.6 MB, less than 58.11% of Java online submissions for Intersection of Two Arrays.
**/

class Solution {
    public int[] intersection(int[] nums1, int[] nums2) {
	  HashSet<Integer> nums2Set = new HashSet<Integer>();
	  HashSet<Integer> resultSet = new HashSet<Integer>();
	  
	  for(int j =0;j<nums2.length;j++)
		  nums2Set.add(nums2[j]);
	  for(int num :nums1){
		  if(nums2Set.contains(num)){
			resultSet.add(num);
		  }
	  }
      int[] result = new int[resultSet.size()];
	  int c =0;
	  for(int n : resultSet)
		  result[c++]=n;
	  
	  return result;
    }
}