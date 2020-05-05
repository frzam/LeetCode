func firstUniqChar(s string) int {
    // Create an array of 26 char and count characters out of the string into the array.
    count := make([]int, 26)
    for i:=0;i<len(s);i++{
        count[s[i]-'a']++
    }
    for i:=0;i<len(s);i++{
        if count[s[i]-'a'] == 1{
            return i
        }
    }
    return -1
}
