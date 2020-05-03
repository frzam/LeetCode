func numEquivDominoPairs(dominoes [][]int) int {
    count := make(map[int]int)
    n := 0
    for i:=0;i<len(dominoes);i++{
        mn := min(dominoes[i][0],dominoes[i][1])* 10 
        mx := max(dominoes[i][0],dominoes[i][1])
        if v, ok := count[mn+mx];ok{
            n += v
        }
        count[mn+ mx]++
    }
     return n
}

func max(a, b int)int{
    if a>b{
        return a
    }
    return b
}
func min(a, b int)int{
    if a <b{
        return a
    }
    return b
}
