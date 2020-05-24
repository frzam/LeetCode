/**
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
