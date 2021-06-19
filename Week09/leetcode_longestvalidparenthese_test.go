package Week09

import (
	"container/list"
	"testing"
)

// dp dp[i] 以i结尾的括号长度 dp[i]=dp[i-1]+dp[i-dp[i-1]-2]+2
func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	maxLength := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			continue
		}
		cur := 0
		if i > 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			cur += dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				cur += dp[i-dp[i-1]-2]
			}
		}
		dp[i] = cur
		if cur > maxLength {
			maxLength = cur
		}

	}
	return maxLength
}

// 栈 维护上一次分割两组括号的pos
func longestValidParentheses2(s string) int {
	q := list.New()
	q.PushBack(-1)
	maxLength := 0
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	for i, c := range s {
		if c == '(' {
			q.PushFront(i)
		} else {
			top := q.Front().Value.(int)
			if top > -1 && s[top] == '(' {
				q.Remove(q.Front())
				maxLength = max(maxLength, i-q.Front().Value.(int))
			} else {
				q.PushFront(i)
			}
		}
	}
	return maxLength
}
// 左右两次循环 -> 实质上是贪心 假如左->右 right>left 说明找到了分割点 同样的 假如右->左 left>right 说明找到了分割点
func longestValidParentheses3(s string) int {
	left,right:=0,0
	maxLength:=0
	max:=func(x,y int) int{
		if x>=y{
			return x
		}
		return y
	}
	for _,c:=range s{
		if c=='('{
			left++
		}else{
			right++
		}
		if left==right{
			maxLength=max(maxLength,left*2)
		}
		if left<right{
			left,right=0,0
		}
	}
	left,right=0,0
	for i:=len(s)-1;i>=0;i--{
		if s[i]==')'{
			right++
		}else{
			left++
		}
		if left==right{
			maxLength=max(maxLength,left*2)
		}
		if left>right{
			left,right=0,0
		}
	}
	return maxLength

}

func TestLong(t *testing.T) {
	longestValidParentheses(
		")()())")
}
