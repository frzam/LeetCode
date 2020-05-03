func maxDistToClosest(seats []int) int {
    last, res, n := -1, 0, len(seats)
    for i:= 0;i<n;i++{
        if seats[i] == 1{
            if last < 0{
                res = i
            }else{
                res = max(res, (i- last)/2)
            }
            last = i

        }
    }
    return max(res, n - last -1)
}

func max(a, b int)int{
    if a > b{
        return a
    }
    return b
}
