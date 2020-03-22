func sortString(s string) string {
    arr := make([]int, 27)
    l := len(s)
    for i:=0;i<len(s);i++{
        arr[int(s[i])-96]++
    }
    var r string
    for l>0{
        for j:= 1;j<len(arr);j++{
            if arr[j] > 0{
                r += string(j+96)
                arr[j] --
                l --
            }
        }
        for j := len(arr)-1;j>0;j--{
            if arr[j]>0{
                r += string(j+96)
                arr[j] --
                l --       
            }
        }
    }
    return r
}