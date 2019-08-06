/**
922. Sort Array By Parity II
Easy

320

32

Favorite

Share
Given an array A of non-negative integers, half of the integers in A are odd, and half of the integers are even.

Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even.

You may return any answer array that satisfies this condition.

 

Example 1:

Input: [4,2,5,7]
Output: [4,5,2,7]
Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.

Solution Analysis:
Runtime: 2 ms, faster than 99.72% of Java online submissions for Sort Array By Parity II.
Memory Usage: 40.1 MB, less than 100.00% of Java online submissions for Sort Array By Parity II.
**/
class Solution {
    public int[] sortArrayByParityII(int[] A) {
        int[] B = new int [A.length];
		int odd = 1;
		int even = 0;
		for (int i =0;i<A.length;i++){
			if(A[i]%2 == 0){
				B[even] = A[i];
				even = even + 2;
			}
			else{
				B[odd] = A[i];
				odd = odd +2;
			}
		}
		return B;
    }
}