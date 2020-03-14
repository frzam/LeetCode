func fairCandySwap(A []int, B []int) []int {
    sumA := 0
    sumB := 0
    for i:=0;i<len(A);i++{
        sumA += A[i]
    }
    for j:=0;j<len(B);j++{
        sumB += B[j]
    }
    diff := (sumA-sumB)/2
    res := make([]int,2)
    for i:=0;i<len(A);i++{
        for j:=0;j<len(B);j++{
            if A[i]-B[j] == diff{
                res[0] = A[i]
                res[1] = B[j]
                return res
            }
        }
    }
    return res
}