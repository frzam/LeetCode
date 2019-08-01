/**
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/


//1) Make an array of nums1+nums2 length.
//2) Start a while loop and compare both nums1 and nums2 array and which ever is lesser assign to new array.
//3) Assign all those elements that are left.
//4) Find the median of new array.

class Solution {
    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int nums[] = new int[nums1.length+nums2.length];
		int i =0;
		int j =0;
		int k =0;
		while(i<nums1.length && j<nums2.length){
			if(nums1[i]<nums2[j])
				nums[k++] = nums1[i++];
			else
				nums[k++] = nums2[j++];
		}
		while(i<nums1.length)
			nums[k++] = nums1[i++];
		while(j<nums2.length)
			nums[k++] = nums2[j++];
		double med = 0;
		int len = nums.length;
		if (len%2 == 0)
			med = (nums[(len-1)/2]+nums[len/2])/2.0;
		else
			med = nums[(len-1)/2];
		return med;
    }
}
