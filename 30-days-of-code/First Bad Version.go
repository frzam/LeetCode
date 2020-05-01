/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    if n < 2{
        return n
    }
   return binarySearch(1,n)
}

func binarySearch(l, r int)int{
    mid := (l+r)/2
    if isBadVersion(mid) && !isBadVersion(mid-1){
        return mid
    }else if isBadVersion(mid){
        return binarySearch(l,mid-1)
    }else{
        return binarySearch(mid+1,r)
    }
    return -1
}
