func distributeCandies(candyType []int) int {
    count := make(map[int]struct{})
    n := len(candyType)/2
    for i := 0;i<len(candyType);i++{
        count[candyType[i]]= struct{}{}
    }
    if n <= len(count){
        return n
    }
    return len(count)
}
