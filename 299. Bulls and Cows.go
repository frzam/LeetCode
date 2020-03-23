func getHint(secret string, guess string) string {
    a := 0
    b := 0
    
    count := make([]int,10)
    
    for i:= 0;i<len(secret);i++{
        if int(secret[i]) == int(guess[i]){
            a++
        }else{
            n := int(secret[i]-'0')
           
            if (count[n])< 0{
                b++
            }
            count[n]++
            m := int(guess[i]-'0')
           
            if (count[m]) >0{
                b++
            }
             count[m]--
        }
    }

   return fmt.Sprintf("%dA%dB",a,b)
}