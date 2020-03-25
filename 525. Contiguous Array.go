func findMaxLength(nums []int) int {
    for i:=0 ;i<len(nums);i++{
        if nums[i]== 0{
            nums[i] = -1
        }
    }
    max := 0
    sum := 0
    count := make(map[int]int)
    count[0]=-1
    for i:= 0;i<len(nums);i++{
        sum += nums[i]
        if n, ok := count[sum];ok{
            if max < i -n{
                max = i -n
            } 
        }else{
            count[sum]= i
        }
    }
    return max
}