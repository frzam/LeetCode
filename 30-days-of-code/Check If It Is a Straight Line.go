func checkStraightLine(coordinates [][]int) bool {
    start := coordinates[0]
    end := coordinates[len(coordinates)-1]
    m := 0
    if end[0] - start[0] != 0{
        m = (end[1] - start[1])/ (end[0]-start[0])
    }
    c := start[1] - m * start[0]
    for i:=1 ;i<len(coordinates)-1;i++{
        if coordinates[i][1] != m * coordinates[i][0] +c{
            return false
        }
    }
    return true
    
}
