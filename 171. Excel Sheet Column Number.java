/**
Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28 
    ...
Example 1:

Input: "A"
Output: 1
Example 2:

Input: "AB"
Output: 28
Example 3:

Input: "ZY"
Output: 701

Solution Analysis:
Runtime: 1 ms, faster than 100.00% of Java online submissions for Excel Sheet Column Number.
Memory Usage: 36.2 MB, less than 100.00% of Java online submissions for Excel Sheet Column Number.

**/

class Solution {
    public int titleToNumber(String s) {
        int num = 0;
        for (int i=0;i<s.length();i++){
            num = ((int)s.charAt(i) - 64) + num *26;
        }
        return num;
    }
}