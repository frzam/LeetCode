/*
Algorithm :
r is used to denote '('
count : it counts total no. of not closed brakets.

1. Iterate through each character and if found ')' and we have no right bracket'(', the add count ++.
2. Else if found ')' and we have r > 0 that we means we have one complete bracket (), so we will decrease the value of r.
3. Else if We have '(', then we will simply increase the value of r.
4. At last we will add all the current count and right '(' that are left.
*/

func minAddToMakeValid(S string) int {
    r := 0
    count := 0
    
    for i:= 0;i<len(S);i++{
        if S[i] == ')' && r == 0{
            count ++
        }else if S[i] == ')' && r > 0{
            r --
        }else if S[i] == '(' {
            r ++
        }
    }
    count += r
    return count
}