package Week06_07

//1. dfs + 剪枝
func generateParenthesis1(n int) []string {
	res:=make([]string,0)
	helper(0,0,n,"",&res)
	return res

}
func helper(left,right,n int,prev string,res *[]string){
	if left==n && right==n{
		*res=append(*res,prev)
		return
	}
	if left<n{
		helper(left+1,right,n,prev+"(",res)
	}
	if right<left{
		helper(left,right+1,n,prev+")",res)
	}
}
// 2. dp
func generateParenthesis2(n int) []string {
	dp:=make([][]string,n+1,n+1)
	dp[0]=[]string{""}
	dp[1]=[]string{"()"}
	for i:=2;i<=n;i++{
		for j:=0;j<i;j++{
			former:=dp[j]
			latter:=dp[i-j-1]
			for _,f:=range former{
				for _,l:=range latter{
					dp[i]=append(dp[i],"("+f+")"+l)
				}
			}
		}
	}
	return dp[n]

}