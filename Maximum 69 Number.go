func maximum69Number (num int) int {
    loc := -1
    n := num
    r := 0
    for i:= 0;n > 0;i++{
        r = n % 10
        if r == 6{
            loc = i
        }
        n = n /10
    }
    if loc == -1{
        return num
    }
    return num + 3 * int(math.Pow(10,float64(loc)))
}
