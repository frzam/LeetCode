/**
In a string S of lowercase letters, these letters form consecutive groups of the same character.

For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and "yy".

Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.

The final answer should be in lexicographic order.

 

Example 1:

Input: "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.
Example 2:

Input: "abc"
Output: []
Explanation: We have "a","b" and "c" but no large group.
Example 3:

Input: "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]
Solution Analysis:
Runtime: 1 ms, faster than 100.00% of Java online submissions for Positions of Large Groups.
Memory Usage: 37.6 MB, less than 94.74% of Java online submissions for Positions of Large Groups.
**/

class Solution {
    public List<List<Integer>> largeGroupPositions(String S) {
        List<List<Integer>> result = new ArrayList();
        int start =0;
      
        for(int i =1;i<=S.length();i++){
           if(i == S.length() || S.charAt(start) != S.charAt(i)){
               if(i-start >=3 )
                  result.add(Arrays.asList(start,i-1));
               start = i;
           }
        }
        return result;
    }
}