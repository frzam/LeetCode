/**
Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.

 

Example 1:

Input: date = "2019-01-09"
Output: 9
Explanation: Given date is the 9th day of the year in 2019.
Example 2:

Input: date = "2019-02-10"
Output: 41
Example 3:

Input: date = "2003-03-01"
Output: 60
Example 4:

Input: date = "2004-03-01"
Output: 61

Solution Analysis:
Runtime: 1 ms, faster than 100.00% of Java online submissions for Day of the Year.
Memory Usage: 34.6 MB, less than 100.00% of Java online submissions for Day of the Year.
**/
class Solution {
    public int dayOfYear(String date) {
        int days =0;
        int year = Integer.parseInt(date.substring(0,4));
        int month = Integer.parseInt(date.substring(5,7));
        int day = Integer.parseInt(date.substring(8,10));
        int feb = 28;
        if(year%400 == 0 ||(year %100 != 0 && year % 4 == 0))
            feb ++;
        int[] months = {31,feb,31,30,31,30,31,31,30,31,30,31};
        for(int i = 0; i<month-1;i++){
            days +=months[i];
        }
        return days+day;
    }
}