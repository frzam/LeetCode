func canConstruct(ransomNote string, magazine string) bool {
    count := make([]int, 26)
    for i:=0;i<len(magazine);i++{
        count[magazine[i]- 'a']++
    }
    for j:= 0;j<len(ransomNote);j++{
        count[ransomNote[j]- 'a']--
        if count[ransomNote[j]-'a'] <0{
            return false
        }
    }
    return true
}
