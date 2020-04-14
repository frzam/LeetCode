func sumZero(n int) []int {
    nums := make([]int, n)
    l := n
    if n %2 != 0{
        l = n -1
    }
    j := 1
    c := 0
    for i:=0;i<l;i++{
        if c % 2 ==0{
            nums[i]= j
            c ++
        }else {
            nums[i] = -j
            c++
            j++
        }
    }
    if l <n{
        nums[n-1]= 0
    }
    return nums
}
