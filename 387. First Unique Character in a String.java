/**
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.

Solution Analysis:
Runtime: 8 ms, faster than 83.79% of Java online submissions for First Unique Character in a String.
Memory Usage: 36.8 MB, less than 100.00% of Java online submissions for First Unique Character in a String.
**/



class Solution {
    public int firstUniqChar(String s) {
     int[] countArr = new int[26];
        for(int i=0;i<s.length();i++)
            countArr[s.charAt(i)-'a']++;
        for(int j=0;j<s.length();j++){
            if(countArr[s.charAt(j)-'a'] == 1)
                return j;
        }
        return -1;
    }
}