/**
In a array A of size 2N, there are N+1 unique elements, and exactly one of these elements is repeated N times.

Return the element repeated N times.

 

Example 1:

Input: [1,2,3,3]
Output: 3
Example 2:

Input: [2,1,2,5,3,2]
Output: 2
Example 3:

Input: [5,1,5,2,5,3,5,4]
Output: 5
 

Note:

4 <= A.length <= 10000
0 <= A[i] < 10000
A.length is even

Solution Analysis:
Runtime: 0 ms, faster than 100.00% of Java online submissions for N-Repeated Element in Size 2N Array.
Memory Usage: 38.5 MB, less than 99.07% of Java online submissions for N-Repeated Element in Size 2N Array.
**/

class Solution {
    public int repeatedNTimes(int[] A) {
		ArrayList<Integer> list = new ArrayList<>();
		for(int i = 0; i<A.length;i++){
			if(!list.contains(A[i]))
				list.add(A[i]);
			else
				return A[i];
		}
	 return -1;
    }
}