/**
Given an input string, reverse the string word by word.

 

Example 1:

Input: "the sky is blue"
Output: "blue is sky the"
Example 2:

Input: "  hello world!  "
Output: "world! hello"
Explanation: Your reversed string should not contain leading or trailing spaces.
Example 3:

Input: "a good   example"
Output: "example good a"
Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
 

Note:

A word is defined as a sequence of non-space characters.
Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
You need to reduce multiple spaces between two words to a single space in the reversed string.
 
**/
class Solution {
    public String reverseWords(String s) {
        String reverseWord = "";
		boolean space = false;
		String word = "";
		if(s.length()==0)
			return "";
		for(int i=0;i<s.length();i++){
			char ch = s.charAt(i);

			if(ch==' ' && !word.equals("") ){
				space = true;
				continue;
			}
			if(!word.equals("") && space==true && reverseWord.equals(""))
			{
				reverseWord = word;
				word = "";
				space=false;
			}
			if(!word.equals("") && space==true){
				reverseWord = word +" "+reverseWord;
				word = "";
				space=false;
			}
			
			if(ch != ' ')
				word = word+ch;


		}
		reverseWord = word +" "+reverseWord;
		reverseWord=reverseWord.trim();
		return reverseWord;
    }
}