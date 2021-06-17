package Week10

// 贪心 如果s[l]==s[r] 删除的字符一定在[l+1,r-1] -> 时间复杂度 ：O(N) 不会重复访问任何一个字符
func validPalindrome(s string) bool {

	return isValid([]rune(s),0,len(s)-1,true)

}
func isValid(cs []rune,l,r int,canDelete bool) bool{
	for l<r{
		if cs[l]!=cs[r] && canDelete{
			return isValid(cs,l+1,r,false) || isValid(cs,l,r-1,false)
		}
		if cs[l]!=cs[r]{
			return false
		}
		return isValid(cs,l+1,r-1,canDelete)
	}
	return true
}