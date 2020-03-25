func isIsomorphic(s string, t string) bool {
    count := make(map[byte]byte)
    for i:=0;i<len(s);i++{
        c, ok := count[s[i]]
        if ok{
            if c == t[i]{
              continue  
            }else{
            return false
            }
        }else{
            for _,v := range count{
                if v == t[i]{
                    return false
                }
            }
            count[s[i]] = t[i]
        }
    }
    return true
}