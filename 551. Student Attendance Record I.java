
/**
You are given a string representing an attendance record for a student. The record only contains the following three characters:
'A' : Absent.
'L' : Late.
'P' : Present.
A student could be rewarded if his attendance record doesn't contain more than one 'A' (absent) or more than two continuous 'L' (late).

You need to return whether the student could be rewarded according to his attendance record.

Example 1:
Input: "PPALLP"
Output: True
Example 2:
Input: "PPALLL"
Output: False

Algorithm:
Loop through all the characters in the string and add absent count if you find 'A'.
If find 'L' check i-1 and i+1 to see if it contains 'L', then return false,
If absent is greater than 1 then return false else true.

**/

class Solution {
    public boolean checkRecord(String s) {
       int absent = 0;
       for(int i =0;i<s.length();i++){
           if(s.charAt(i) == 'A')
               absent++;
           if(i>0 && i <s.length()-1 && s.charAt(i)=='L'){
             if(s.charAt(i-1)=='L'&&s.charAt(i+1)=='L')
                 return false;
           }
       } 
        if(absent >1)
            return false;
    
    return true;
    }
}