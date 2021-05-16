package Week06_07

// 1.dp
func countSubstrings1(s string) int {
	n := len(s)
	dps := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		dps[i] = make([]bool, n, n)
	}
	cs := []rune(s)
	count := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if j == i {
				dps[i][j] = true
				count++
				continue
			}
			if j == i+1 {
				if cs[i] == cs[j] {
					dps[i][j] = true
					count++
				}
				continue
			}
			dps[i][j] = dps[i+1][j-1] && cs[i] == cs[j]
			if dps[i][j] {
				count++
			}
		}
	}
	return count
}

// 2. 中心扩展
func countSubstrings(s string) int {
	// 1. dp
	// 2. 中心判断向两边扩展
	cs := []rune(s)
	count, n := 0, len(cs)
	c := func(l, r int) int {
		res := 0
		for l >= 0 && r < n && cs[l] == cs[r] {
			res++
			l--
			r++
		}
		return res
	}

	for i := 0; i < n; i++ {
		count += c(i, i)
		count += c(i, i+1)
	}
	return count

}
