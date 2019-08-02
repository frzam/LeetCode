/**
Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written.

Do the above modifications to the input array in place, do not return anything from your function.

Runtime Analysis:
Runtime: 0 ms, faster than 100.00% of Java online submissions for Duplicate Zeros.
Memory Usage: 37.5 MB, less than 100.00% of Java online submissions for Duplicate Zeros.
**/

class Solution {
    public void duplicateZeros(int[] arr) {
		int len = arr.length;
		int[] copyArr = arr.clone();
		int index = 0;
		for(int i=0; i<len;i++){
			if (index>= len)
				break;
			else if(copyArr[i]==0 && index+1<len){
				arr[index++] = 0;
				arr[index++] = 0;
				}
			else{
					arr[index++] = copyArr[i];
			}
				
		}
    }
}