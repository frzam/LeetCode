func findAndReplacePattern(words []string, pattern string) []string {
    pTable := make(map[byte]int)
    var res []string
    for i:=0;i<len(pattern);i++{
        pTable[pattern[i]] ++
    }
    for i:=0;i<len(words);i++{
        wCount := make(map[byte]int)
        for j:=0;j<len(words[i]); j++{
            wCount[words[i][j]]++
        }
       
        if len(wCount) == len(pTable){
            cTable := make(map[byte]byte)
            count := 0
            for j:=0;j<len(words[i]);j++{
                if v, ok := cTable[words[i][j]]; ok {
                    if v != pattern[j]{
                        break
                    }else {
                        count ++
                    }
                 
                }else{                    
                    count ++
                    cTable[words[i][j]] = pattern[j]
                }
            }
            if count == len(pattern){
                res = append(res, words[i])
            }
        }
    }
    return res
}
