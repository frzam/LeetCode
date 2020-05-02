func diStringMatch(S string) []int {
    res := make([]int, len(S)+1)
    h := len(S)
    l := 0
    for i:=0;i<len(S);i++{
        if S[i] == 'I'{
            res[i] = l
            l++
        }else{
            res[i] = h
            h --
        }
    }
    res[len(S)]= h
    return res
}
