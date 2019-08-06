/**
Students are asked to stand in non-decreasing order of heights for an annual photo.

Return the minimum number of students not standing in the right positions.  (This is the number of students that must move in order for all students to be standing in non-decreasing order of height.)

 

Example 1:

Input: [1,1,4,2,1,3]
Output: 3
Explanation: 
Students with heights 4, 3 and the last 1 are not standing in the right positions.
 

Note:

1 <= heights.length <= 100
1 <= heights[i] <= 100

Runtime: 0 ms, faster than 100.00% of Java online submissions for Height Checker.
Memory Usage: 34.4 MB, less than 100.00% of Java online submissions for Height Checker
**/

class Solution {
    public int heightChecker(int[] heights) {
       int[]  heightFreq = new int[101];
	   int currentHeight = 0;
	   int count =0;
	    for (int height : heights)
			heightFreq[height] ++;
		
		for(int i = 0;i<heights.length;i++){
			while (heightFreq[currentHeight] == 0)
				currentHeight ++;
			if(currentHeight != heights[i])
				count ++;
			heightFreq[currentHeight] --;
		}
		return count;
    }
}