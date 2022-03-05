package brush

import (
	"sort"
	"testing"
)

func nextGreaterElement(n int) int {
	target:=make([]int,0)
	t:=n
	for t>0{
		target=append(target,t%10)
		t/=10
	}
	reverse:=func(nums []int){
		l,r:=0,len(nums)-1
		for l<r{
			nums[l],nums[r]=nums[r],nums[l]
			l++
			r--
		}
	}
	reverse(target)
	for i:=len(target)-2;i>=0;i--{
		if target[i]<target[i+1]{
			// 正序
			to:=i+1
			for j:=i+2;j<len(target);j++{
				if target[j]>target[i] && target[j]<target[to]{
					to=j
				}
			}
			target[to],target[i]=target[i],target[to]
			tt:=target[i+1:]
			sort.Slice(tt,func(x,y int) bool{
				return target[x]<target[y]
			})
			break
		}
	}

	left:=0
	for _,n:=range target{
		left*=10
		left+=n
	}
	if left==n{
		return -1
	}
	return left
}

func TestNextGreaterNumber(t *testing.T){
	nextGreaterElement(1200000)
}
