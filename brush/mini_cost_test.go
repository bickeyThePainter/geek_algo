package brush

import (
	"fmt"
	"math"
	"testing"
)

func mincostTickets(days []int, costs []int) int {
	mem:=make([]int,366)
	dayIndex:=make(map[int]bool)
	for _,d:=range days{
		dayIndex[d]=true
	}
	min:=func(x,y int) int{
		if x==-1{
			return y
		}
		if x<y{
			return x
		}
		return y
	}
	costIndex:=make(map[int]int)
	costIndex[1]=costs[0]
	costIndex[7]=costs[1]
	costIndex[30]=costs[2]
	for i:=365;i>=1;i--{
		if !dayIndex[i]{
			continue
		}
		cur:=-1
		for d,c:=range costIndex{
			//if i+d<365{
			cur=min(cur,c+func() int{
				if i+d>365{
					return 0
				}
				return mem[i+d]
			}())
			//}
		}
		mem[i]=cur
	}
	return mem[1]
}

func TestMiniCost(t *testing.T){
	//mincostTickets([]int{1,2,3,4,5,6,7,8,9,10,30,31},[]int{2,7,15})
	//longestValidParentheses(")()()())")
	superEggDrop(3,11)
	/**
	[bs1] l=1,r=4,mid=2,egg=2
	[bs1] l=1,r=2,mid=1,egg=2
	[bs] l=1,r=4,mid=2,egg=2
	[bs] l=3,r=4,mid=3,egg=2
	 */
}

func superEggDrop(k int, n int) int {
	// 朴素
	dp:=make([][]int,n+1)
	for i:=0;i<=n;i++{
		dp[i]=make([]int,k+1)
		if i>0{
			dp[i][0]=math.MaxInt32
		}
	}
	plusOne:=func(x int) int{
		if x==math.MaxInt32{
			return x
		}
		return x+1
	}
	for i:=1;i<=n;i++{
		for j:=1;j<=k;j++{
			dp[i][j]=plusOne(bs1(dp,j,i))
		}
	}
	return dp[n][k]
}

func bs(dp [][]int,egg,level int) int{
	l,r:=1,level
	for l<r{
		mid:=l+(r-l)>>1
		if dp[mid-1][egg-1]<dp[level-mid][egg]{
			l=mid+1
		}else{
			r=mid
		}
	}
	if dp[l-1][egg-1]==dp[level-l][egg]{
		return dp[l-1][egg-1]
	}
	min:=func(x, y int) int{
		if x<y{
			return x
		}
		return y
	}
	max:=func(x,y int) int{
		if x>y{
			return x
		}
		return y
	}
	return min(max(dp[l-1][egg-1],dp[level-l][egg]),max(dp[l-2][egg-1],dp[level-l+1][egg]))
}

func bs1(dp [][]int,egg,level int) int{
	max:=func(x,y int) int{
		if x>y{
			return x
		}
		return y
	}
	l,r:=1,level
	min:=func(x,y int) int{
		if x<y{
			return x
		}
		return y
	}
	realMin:=math.MaxInt32
	for i:=1;i<=r;i++{
		fmt.Printf("[bs1][%d-%d][%d] seq i=%d,calc=%d\n",l,r,egg,i,max(dp[i-1][egg-1],dp[level-i][egg]))
		realMin=min(realMin,max(dp[i-1][egg-1],dp[level-i][egg]))
	}
    lt,rt:=l,r
	for l<r{
		mid:=l+(r-l)>>1
		if max(dp[mid-1][egg-1],dp[level-mid][egg])>max(dp[mid][egg-1],dp[level-mid-1][egg]){
			l=mid+1
		}else{
			r=mid
		}
	}
	if realMin!=max(dp[l-1][egg-1],dp[level-l][egg]){
		fmt.Printf("[bs1][%d-%d][%d] wrong answer calc=%d,should=%d\n",lt,rt,egg,max(dp[l-1][egg-1],dp[level-1][egg]),realMin)
	}

	return max(dp[l-1][egg-1],dp[level-l][egg])
}
/**
[bs1] l=1,r=4,mid=2,egg=2
[bs1] dp[mid-1][egg-1]=2,dp[level-mid][egg]=2,dp[mid][egg-1]=2,dp[level-mid-1][egg]=1
[bs1] l=1,r=2,mid=1,egg=2
[bs1] dp[mid-1][egg-1]=1,dp[level-mid][egg]=2,dp[mid][egg-1]=1,dp[level-mid-1][egg]=2
[bs] l=1,r=4,mid=2,egg=2
[bs] dp[mid-1][egg-1]=2,dp[level-mid][egg]=2
[bs] l=3,r=4,mid=3,egg=2
[bs] dp[mid-1][egg-1]=3,dp[level-mid][egg]=1
 */

/*
[bs1][1-7] seq i=1,calc=6
[bs1][1-7] seq i=2,calc=2147483647
[bs1][1-7] seq i=3,calc=2147483647
[bs1][1-7] seq i=4,calc=2147483647
[bs1][1-7] seq i=5,calc=2147483647
[bs1][1-7] seq i=6,calc=2147483647
[bs1][1-7] seq i=7,calc=2147483647
[bs1][1-7] seq i=1,calc=3
[bs1][1-7] seq i=2,calc=3
[bs1][1-7] seq i=3,calc=3
[bs1][1-7] seq i=4,calc=3
[bs1][1-7] seq i=5,calc=4
[bs1][1-7] seq i=6,calc=5
[bs1][1-7] seq i=7,calc=6
[bs1][1-7] seq i=1,calc=3
[bs1][1-7] seq i=2,calc=3
[bs1][1-7] seq i=3,calc=3
[bs1][1-7] seq i=4,calc=2
[bs1][1-7] seq i=5,calc=3
[bs1][1-7] seq i=6,calc=3
[bs1][1-7] seq i=7,calc=3
[bs1][1-7] wrong answer calc=3,should=2
 */



