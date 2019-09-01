/**
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5

Solution Analysis:
Runtime: 0 ms, faster than 100.00% of Java online submissions for Length of Last Word.
Memory Usage: 35.9 MB, less than 100.00% of Java online submissions for Length of Last Word.
**/

/**
Algorithm:
1. Return 0 if the length is 0.
2. Loop through the string from back side and when find a char other than space, make found as true and count ++.
3. If found is true and we get space that means we have completed the world, so just return the count. Ex.  ' 'World, so return 5.
4. If we don't have any space in the string then we will return count when we move to first char.
**/

class Solution {
    public int lengthOfLastWord(String s) {
        int count =0;
        boolean found = false;
        
        if(s.length()==0)
            return 0;
        for(int i=s.length()-1;i>=0;i--){
            if(s.charAt(i) != ' '){
                found = true;
                count ++;
            }
            if(found == true && s.charAt(i)== ' ')
                return count;
            if(found == true && i == 0)
                return count;
        }
        return 0;
    }
}