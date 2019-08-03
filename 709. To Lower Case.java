/**
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

 

Example 1:

Input: "Hello"
Output: "hello"
Example 2:

Input: "here"
Output: "here"
Example 3:

Input: "LOVELY"
Output: "lovely"

Solution Analysis:
Runtime: 0 ms, faster than 100.00% of Java online submissions for To Lower Case.
Memory Usage: 33.9 MB, less than 99.75% of Java online submissions for To Lower Case.
**/

class Solution {
    public String toLowerCase(String str) {
        String lowerCaseStr = "";
        for(int i =0;i<str.length();i++){
            if((int)str.charAt(i) >=65 && (int)str.charAt(i) <=90)
               lowerCaseStr += (char)((int)str.charAt(i)+32);
            else
                lowerCaseStr += str.charAt(i);
        }
        return lowerCaseStr;
    }
}