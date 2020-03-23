func sortColors(nums []int)  {
    low := 0
    high := len(nums)-1
    
    for i := low;i<=high;{
        if nums[i] == 2{
               nums[high],nums[i]= nums[i],nums[high]
                high --
        }else if nums[i] == 0{
                nums[low],nums[i]= nums[i],nums[low]
                i++
                low ++
        }else{
            i++
        }
        
    }
}