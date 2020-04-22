func uniqueOccurrences(arr []int) bool {
    count := make(map[int]int)
    for i:=0;i<len(arr);i++{
        count[arr[i]] ++
    }
    unique := make(map[int]struct{})
    for _,v := range count{
        if _, ok := unique[v];ok{
            return false
        }else{
            unique[v] = struct{}{}
        }
    } 
    return true
}
