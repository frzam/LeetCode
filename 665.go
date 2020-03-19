func checkPossibility(nums []int) bool {
    count := 0
    for i:=0 ;i < len(nums)-1 ;i++{
        if nums[i]>nums[i+1]{
            if count > 0{
                return false
            }
            if i - 1 >= 0 && i + 2 < len(nums)&& (nums[i] > nums[i + 2] && nums[i + 1] < nums[i - 1]){
            return false
     
        }
            count ++
        }
       
    }
return true
}