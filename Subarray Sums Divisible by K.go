func subarraysDivByK(A []int, K int) int {
    m := make([]int, K)
    m[0] = 1
    sum := 0
    count :=0
    for _,a := range A{
        sum = (sum +a)%K
       
        if sum <0{
            sum += K
        }
        count += m[sum]
        m[sum]++
    }
    return count
}
