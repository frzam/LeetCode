func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    count := 0
    dont := false
    for i:= 0;i<len(arr1);i++{
        dont = false
        for j:= 0;j<len(arr2);j++{
            if mod(arr1[i] - arr2[j]) <= d{
                dont = true
                break
            }
        }
        if !dont{
         count ++   
        }
       }
    return count
}
func mod(n int)int{
    if n < 0{
        return n * -1
    }
    return n
}
