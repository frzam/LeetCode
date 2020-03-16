package main

import "fmt"
func containsNearbyDuplicate(nums []int, k int) bool {
    m := make(map[int]int)
	for i:=0;i<len(nums);i++{
		m[nums[i]] = m[nums[i]]+1
	}
	for key,v := range m{
		if v > 1{
			index := indexDiff(nums,key)
			if index <= k{
			
				return true
			}
		}
	}
	return false
}

func indexDiff(nums []int,k int) int{
	i := -1
	diff := len(nums)
	for j:= 0;j<len(nums);j++{
		if k == nums[j] && i == -1{
			i = j
		}else if (k == nums[j]){
			if j-i < diff{
			diff = j-i
			i = j
		}
		}
	}
	
	return diff
}

func main(){
nums := []int {1,2,3,1,2,3}
res := containsNearbyDuplicate(nums,2)
fmt.Println(res)
}