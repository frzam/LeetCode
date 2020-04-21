func isMonotonic(A []int) bool {
    
     up := false
     down := false
     for i:=0;i<len(A)-1;i++{
         if A[i]< A[i+1]{
             up = true
         }else if A[i]>A[i+1]{
             down = true
         }else if up && down{
             return false
         }
    }
    return !(up && down)
}
