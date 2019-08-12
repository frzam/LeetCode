/**
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

Runtime: 2 ms, faster than 94.32% of Java online submissions for Container With Most Water.
Memory Usage: 40.1 MB, less than 93.59% of Java online submissions for Container With Most Water.
**/
class Solution {
    public int maxArea(int[] height) {
		int min = height[0];
		int max = 0;
		int i=0; int j = height.length-1;
		while(i<j){
			if(height[i]<height[j])
				min = height[i];
			else 
				min = height[j];
			if(max < min*(j-i))
				max = min*(j-i);
			if(height[i]<height[j])
				i++;
			else
				j--;
		}
		return max;
    }
}
