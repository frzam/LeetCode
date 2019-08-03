/**
Validate if a given string can be interpreted as a decimal number.

Some examples:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
" -90e3   " => true
" 1e" => false
"e3" => false
" 6e-1" => true
" 99e2.5 " => false
"53.5e93" => true
" --6 " => false
"-+3" => false
"95a54e53" => false

Note: It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one. However, here is a list of characters that can be in a valid decimal number:

Numbers 0-9
Exponent - "e"
Positive/negative sign - "+"/"-"
Decimal point - "."
Of course, the context of these characters also matters in the input.

Runtime Analysis:
Runtime: 7 ms, faster than 22.35% of Java online submissions for Valid Number.
Memory Usage: 36.3 MB, less than 98.56% of Java online submissions for Valid Number.
**/
class Solution {
    public boolean isNumber(String s) {
        try{
            double num = Double.parseDouble(s);
            if(s.charAt(s.length()-1)=='f' || s.charAt(s.length()-1)=='D' ) // Because it is taking "959440.94f" not as decimal.
                return false;
        }
        catch (Exception e){
            return false;
        }
        return true;
    }
}