/**
Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

Note:
You may assume that both strings contain only lowercase letters.

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true

Runtime: 4 ms, faster than 72.47% of Java online submissions for Ransom Note.
Memory Usage: 37.6 MB, less than 100.00% of Java online submissions for Ransom Note.

**/
class Solution {
    public boolean canConstruct(String ransomNote, String magazine) {
     int mag[] = new int[26];
        for(int i =0;i<magazine.length();i++)
            mag[magazine.charAt(i)-'a']++;
        for(int j =0;j<ransomNote.length();j++){
             if(--mag[ransomNote.charAt(j)- 'a']<0)
                return false;
        }
          
        return true;        
    }
}