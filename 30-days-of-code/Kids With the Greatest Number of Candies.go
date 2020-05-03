func kidsWithCandies(candies []int, extraCandies int) []bool {
    res := make([]bool, len(candies))
    max :=  0
    for i:=0;i<len(candies);i++{
        if candies[i] > max{
            max = candies[i]
        }
    }
    for j:=0;j<len(candies);j++{
        if candies[j]+extraCandies >=max{
            res[j] = true
        }
    }
    return res
}
