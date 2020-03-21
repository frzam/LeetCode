func merge(nums1 []int, m int, nums2 []int, n int)  {
    i := m -1
    j := n -1
    k := len(nums1) - 1
    for j >= 0 && k>=0 && i >= 0{
        if nums2[j] > nums1[i]{
            nums1[k] = nums2[j]
            j --
        }else{
            nums1[k] = nums1[i]
            i --
        }
        k --
    }
    for i >= 0 && k >= 0{
        nums1[k] = nums1[i]
        i --
        k --
    }
        for j >= 0 && k >= 0{
        nums1[k] = nums2[j]
        j --
        k --
    }
}
