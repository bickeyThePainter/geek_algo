package Week10
// 通配符匹配
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 0; i < m; i++ {
		if p[i] == '*' {
			dp[0][i+1] = dp[0][i]
		}
	}
	for i, s1 := range s {
		for j, p1 := range p {
			if p1 != '?' && p1 != '*' {
				dp[i+1][j+1] = dp[i][j] && s1 == p1
			} else if p1 == '?' {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = dp[i][j] || dp[i+1][j] || dp[i][j+1]
			}
		}
	}
	return dp[n][m]
}
