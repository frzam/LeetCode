func findComplement(num int) int {
    p:= 1
    for p < num{
        p = p * 2 + 1        
    }
    return p^num
}
