func maxSatisfaction(sat []int) int {
    count := make([]int, 2001)
    for i:=0;i<len(sat);i++{
        count[sat[i]+1000]++
    }
    c := 0
    for j:= 0;j<len(count);j++{
        for count[j] >0{
            sat[c] = j - 1000
            count[j]--
            c++
        }
    }
    z := len(sat)-1
    min := sat[z]
    for k:= 0;k<len(sat);k++{
        if sat[k] >=0 && min >sat[k]{
            z = k
            min = sat[k]
        }
    }   
    sc := 1
    mul := 0
    max := 0
    for z >= 0{
        for i:= z;i<len(sat);i++{
            mul += sc * sat[i]
            sc ++
        }
        if mul >= max{
            z --
            sc = 1
            max = mul
            mul =0
        }else{
            return max
        }
    }
    return max
}
