func countElements(arr []int) int {
    countArr := make([]int, 1001)
    for i:= 0;i<len(arr);i++{
        countArr[arr[i]] ++
    }
    count := 0
    for j:=0;j<len(countArr)-1;j++{
        if countArr[j] > 0 && countArr[j+1] >0{
            count += countArr[j]
        }
    }
    return count
}
