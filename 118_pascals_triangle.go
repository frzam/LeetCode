package leetcode

func generate(n int) [][]int {
	var tri [][]int

	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < i+1; j++ {
			if i == 0 || j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, tri[i-1][j]+tri[i-1][j-1])
			}
		}
		tri = append(tri, row)
	}
	return tri
}
