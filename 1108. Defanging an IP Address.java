/***

Example 1:

Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"
Example 2:

Input: address = "255.100.50.0"
Output: "255[.]100[.]50[.]0"

*/
class Solution {
    public String defangIPaddr(String address) {
		String defangedAddress = "";
        for(int i=0; i<address.length();i++){
			if(address.charAt(i) == '.')
				defangedAddress = defangedAddress + "[.]";
			else
				defangedAddress= defangedAddress + address.charAt(i);
			
		}
		return defangedAddress;
    }
}
