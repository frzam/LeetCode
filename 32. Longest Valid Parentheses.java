/**
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"


**/

class Solution {
    public int longestValidParentheses(String s) {
        
		int dp_arr[] = new int[s.length()];
		int longestLength = 0;
		int leftCount = 0;
		
		for(int i = 0; i < s.length(); i++){
			if(s.charAt(i) == '(')
				leftCount ++;
			else if(leftCount > 0){
				dp_arr[i] = dp_arr[i - 1]+2;
				
				if(i - dp_arr[i] >= 0)
					dp_arr[i] = dp_arr[i]+dp_arr[i - dp_arr[i]];
			longestLength = Math.max(longestLength, dp_arr[i]);	
			leftCount --;
			} 
				
		}
    return longestLength;
	}
	
}