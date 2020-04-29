func generateTheString(n int) string {
    var str string
    for i:=0;i<n-1;i++{
        str += "a"
    }
    if n %2 == 0{
        str +="b"
    }else{
        str += "a"
    }
    return str
}
