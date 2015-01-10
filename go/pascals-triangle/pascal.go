package pascal

func Triangle(n int) [][]int {
	triangle := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		triangle = append(triangle, rowAt(i, triangle))
	}
	return triangle
}

func rowAt(n int, triangle [][]int) []int {
	row := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		if i == 0 || i == n {
			row = append(row, 1)
		} else {
			val := valueAt(i, n, triangle)
			row = append(row, val)
		}
	}
	return row
}

func valueAt(x int, y int, triangle [][]int) int {
	prevRow := triangle[y-1]
	return prevRow[x] + prevRow[x-1]
}
