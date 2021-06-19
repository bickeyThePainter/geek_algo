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

func TestLong(t *testing.T) {
	longestValidParentheses(
		")()())")
}
