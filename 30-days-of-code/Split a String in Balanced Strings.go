func balancedStringSplit(s string) int {
    b := 0
    count := 0
    for i:= 0;i<len(s);i++{
        if s[i]== 'L'{
            b ++
        }else{
            b--
        }
        if b == 0{
            count ++
        }
    }
    return count
}
