package Week06_07

// 1. dp
func maximalSquare(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])
	dps := make([][]int, n, n)
	length := 0
	for i := 0; i < n; i++ {
		dps[i] = make([]int, m, m)
	}

	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y

	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dps[i][j] = 1
				} else {
					dps[i][j] = min(min(dps[i][j-1], dps[i-1][j]), dps[i-1][j-1]) + 1
				}
				length = max(length, dps[i][j])
			}
		}
	}
	return length * length
}
