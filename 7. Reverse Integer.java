/**
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.


**/

class Solution {
    public int reverse(int x) {
         long s =0 ; int result =0;
      if(x < Integer.MAX_VALUE && x > Integer.MIN_VALUE){
        int  r = 0;
        while(x != 0){
            r = x %10;
            s = s*10 + r;
            x = x /10;
            
        }
        }
        if(s < Integer.MAX_VALUE && s > Integer.MIN_VALUE){
            result = (int)s;
        }
        return result;
    }
    
}