package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := len(nums1) - 1
	m--
	n--
	for m >= 0 && n >= 0 && i >= 0 {
		if nums1[m] >= nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}
	for n >= 0 && i >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
}
