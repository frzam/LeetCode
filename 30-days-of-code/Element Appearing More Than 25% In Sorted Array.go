func findSpecialInteger(arr []int) int {
    max, count := 1, 1
    num := arr[0]
    for i:= 0;i<len(arr)-1;i++{
        if arr[i] == arr[i+1]{
            count ++
            if count > max{
                max =  count
                num = arr[i]
            }
        }else{
            count = 1
        }
    }
    return num
}
