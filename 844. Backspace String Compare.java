/**
Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

Example 1:

Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
Example 2:

Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
Example 3:

Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
Example 4:

Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
Note:

1 <= S.length <= 200
1 <= T.length <= 200
S and T only contain lowercase letters and '#' characters.
Follow up:

Can you solve it in O(N) time and O(1) space?

Runtime: 1 ms, faster than 73.56% of Java online submissions for Backspace String Compare.
Memory Usage: 34.2 MB, less than 100.00% of Java online submissions for Backspace String Compare.
**/
class Solution {
    public boolean backspaceCompare(String S, String T) {
        ArrayList<Character> sList = new ArrayList<>();
        ArrayList<Character> tList = new ArrayList<>();
        

        for(int i =0;i<S.length();i++){
            if(S.charAt(i)=='#'){
                if(sList.size()>0)
                     sList.remove(sList.size()-1);
             }
            else
                sList.add(S.charAt(i));
        }  
        for(int i =0;i<T.length();i++){
            if(T.charAt(i)=='#'){
                if(tList.size()>0)
                    tList.remove(tList.size()-1);
             }
            else
                tList.add(T.charAt(i));
        }
        if(sList.size() != tList.size())
            return false;
        
        for(int j =0;j<sList.size();j++){
            if(tList.get(j) !=sList.get(j) )
                return false;
        }
        return true;
    }
}