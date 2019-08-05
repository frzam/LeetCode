/**
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.

 

Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Solution Analysis:
Runtime: 1 ms, faster than 100.00% of Java online submissions for Squares of a Sorted Array.
Memory Usage: 40.4 MB, less than 97.35% of Java online submissions for Squares of a Sorted Array.
**/

class Solution {
    public int[] sortedSquares(int[] A) {
    int[] B = new int[A.length];
        int j = 0; int k = A.length-1;
        for(int i =k;i>=0;i--){
            if(Math.abs(A[k]) > Math.abs(A[j])){
                B[i] = A[k] * A[k];
                k--;
            }
            else
            {
                B[i] = A[j] * A[j];
                j++;
            }
        }
        return B;
    }
}