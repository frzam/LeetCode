func subtractProductAndSum(n int) int {
    prod := 1
    r := 0
    s := n
    for n>0 {
        r = n % 10
        prod *= r
        n = n/10
    }
    n = s
    r =0
    sum := 0
    for n>0{
        r =n %10
        sum += r
        n = n/10
    }
    return prod - sum
}
