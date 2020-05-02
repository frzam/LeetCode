func numJewelsInStones(J string, S string) int {
    jewels := make(map[byte]int)
    for i:=0;i<len(J);i++{
        jewels[J[i]] ++
    }
    count := 0
    for i:=0;i<len(S);i++{
        if _,ok := jewels[S[i]];ok{
            count ++
        }
    }
    return count
}
