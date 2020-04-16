func checkIfExist(arr []int) bool {
    check := make(map[int]struct{})
    ok:= false
    for i:=0;i<len(arr);i++{
        if _,ok = check[2*arr[i]];ok{
            return ok
        }
        check[arr[i]] = struct{}{}
    }
    for i:=0;i<len(arr);i++{
        if _,ok = check[2*arr[i]];ok && arr[i] != 0{
            return ok
        }
    }
    return false
}
