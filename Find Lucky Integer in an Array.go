func findLucky(arr []int) int {
    count := make([]int,501)
    for i:= 0;i<len(arr);i++{
        count[arr[i]] ++
    }
    lucky := -1
    for j:=len(count)-1;j>0;j--{
        if j == count[j]{
            lucky = count[j]
            break
        }
    }
    return lucky
}
