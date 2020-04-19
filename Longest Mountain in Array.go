func longestMountain(A []int) int {
    res := 0
    up := 0
    down := 0
    for i:= 1;i<len(A);i++{
        if (down > 0 && A[i-1]<A[i])|| A[i-1] == A[i]{
            up = 0
            down = 0            
        }
        if A[i-1] < A[i]{
            up++
        }
        if A[i-1] > A[i]{
            down ++
        }
        if up >0 && down >0 && up + down + 1 > res{
            res = up + down + 1
        }
    }
    return res
}
