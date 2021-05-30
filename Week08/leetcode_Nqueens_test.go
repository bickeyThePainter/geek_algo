package Week08

// 二进制存储皇后位置

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	prev := make([]int, 0)
	dfs1(0, n, 0, 0, 0, &prev, &res)
	return res
}

func dfs1(index, n, cols, adds, minus int, prev *[]int, res *[][]string) {
	if index == n {
		curRes := make([]string, n, n)
		for i := 0; i < n; i++ {
			tmp := ""
			for j := 0; j < n; j++ {
				if twoStep((*prev)[i]) == j {
					tmp += "Q"
				} else {
					tmp += "."
				}
			}
			curRes[i] = tmp
		}
		*res = append(*res, curRes)
		return
	}
	candidates := (^(cols | adds | minus)) & ((1 << n) - 1)
	for candidates != 0 {
		cur := candidates & (-candidates)
		candidates = (candidates - 1) & candidates
		*prev = append(*prev, cur)
		dfs1(index+1, n, cols|cur, (adds|cur)>>1, (minus|cur)<<1, prev, res)
		*prev = (*prev)[0 : len(*prev)-1]
	}
}
// 可以用dp优化
func twoStep(n int) int {
	count := 0
	for n > 0 {
		n /= 2
		count++
	}
	return count - 1
}
