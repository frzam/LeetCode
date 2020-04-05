func maxProfit(prices []int) int {
    last := 0
    if len(prices) == 0{
        return 0
    }
    secondLast := prices[0]
    for i:=1;i<len(prices);i++{
        last = prices[i] 
        prices[i]= prices[i] - secondLast
        secondLast = last
    }
    sum := 0
    for i:=1 ;i<len(prices);i++{
        if prices[i] >0{
            sum += prices[i]
        }
    }
    return sum
}

	
