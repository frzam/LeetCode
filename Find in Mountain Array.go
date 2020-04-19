/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
    peak := findPeak(mountainArr, 0,mountainArr.length()-1)
    res := find(mountainArr,0,peak,target,"ASC")
    if res == -1{
        return find(mountainArr,peak,mountainArr.length()-1,target,"DESC")
    }
    return res
}
func find(m *MountainArray, l , r, target int,order string)int{
    if l <= r{

    mid := (l+r)/2
    if order == "ASC"{
              
        if m.get(mid)>target{
            return find(m,l,mid-1,target,order)   
        }else if m.get(mid) == target{
            return mid
        }else{
            return find(m,mid+1,r,target,order)
        }

        }else{
         if m.get(mid)>target{
             return find(m,mid+1,r,target,order) 
        }else if m.get(mid) == target{
            return mid
        }else{
             return find(m,l,mid-1,target,order)
        }

        }
        }
    return -1
    }



func findPeak(m *MountainArray, l, r int)int{
    if l == r{
        return l
    }
    mid := (l+r)/2
    if m.get(mid) > m.get(mid+1){
        return findPeak(m,l,mid)
    }else{
        return findPeak(m,mid+1,r)
    }
}
