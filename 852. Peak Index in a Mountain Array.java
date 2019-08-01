/**
Let's call an array A a mountain if the following properties hold:

A.length >= 3
There exists some 0 < i < A.length - 1 such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
Given an array that is definitely a mountain, return any i such that A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1].

Example 1:

Input: [0,1,0]
Output: 1
Example 2:

Input: [0,2,1,0]
Output: 1
Note:

3 <= A.length <= 10000
0 <= A[i] <= 10^6
A is a mountain, as defined above.

**/
/**
Runtime: 0 ms, faster than 100.00% of Java online submissions for Peak Index in a Mountain Array.
Memory Usage: 38.4 MB, less than 100.00% of Java online submissions for Peak Index in a Mountain Array.
*/

//Since it is given that the array is definitely a mountain thus A[i]>A[i+1] must be present only once.
class Solution{
	public int peakIndexInMountainArray(int[] A){
		int i =0;
		for (;i<A.length-1;i++){
			if(A[i]>A[i+1])
				break;
		}
		return i;
	}
}

// If it will not be given that mountain is definitely Mountain then we could have used this solution.


class Solution {
    public int peakIndexInMountainArray(int[] A) {
        int peakLoc = 0;
		int peakFlag = -2;
		
	
	for(int i =0; i<A.length-1;i++){
	 if(i == 0 ){
		if(A[i]>A[i+1]){
            peakLoc = i;
			peakFlag ++;
			} 
	 }	
	 else if(A[i-1]<A[i] && A[i]>A[i+1]){
			peakLoc = i;
			peakFlag ++;
		}
	}
	if(peakFlag == -1)
		return peakLoc;
	else 
		return -1;
    }
}