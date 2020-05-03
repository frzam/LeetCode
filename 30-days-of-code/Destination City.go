func destCity(paths [][]string) string {
    dest := make(map[string]struct{})
    for i:= 0;i<len(paths);i++{
        dest[paths[i][0]]= struct{}{}
        }
    for j:= 0;j<len(paths);j++{
        if _, ok := dest[paths[j][1]];!ok{
            return paths[j][1]
        }
    }
    return ""
}
