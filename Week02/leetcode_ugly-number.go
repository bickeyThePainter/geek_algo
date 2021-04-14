package Week02

import (
	"container/heap"
)

// 1. brute force 超时
// 2. heap 小顶堆+set去重
// 3. three pointers
func nthUglyNumber1(n int) int {
	count := 0
	cur := 0
	for i := 1; count < n; i++ {
		if isUgly(i) {
			cur = i
			count++
		}
	}
	return cur
}

func isUgly(n int) bool {
	if n == 1 {
		return true
	}
	if n%2 == 0 {
		return isUgly(n / 2)
	}

	if n%3 == 0 {
		return isUgly(n / 3)
	}

	if n%5 == 0 {
		return isUgly(n / 5)
	}
	return false

}

func nthUglyNumber2(n int) int {
	h := make(IntHeap, 0)
	heap.Push(&h, 1)
	factors := []int{2, 3, 5}
	check := make(map[int]struct{})
	for i := 1; i < n; i++ {
		top := heap.Pop(&h).(int)
		for _, f := range factors {
			cur := f * top
			if _, ok := check[cur]; !ok {
				check[cur] = struct{}{}
				heap.Push(&h, cur)
			}
		}
	}
	return heap.Pop(&h).(int)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] <= h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func nthUglyNumber3(n int) int {
	factors:=[]int{2,3,5}
	res:=make([]int,n,n)
	i1,i2,i3:=0,0,0
	res[0]=1
	min:=func(x,y,z int) int{
		if x<=y{
			if x<=z{
				return x
			}else{
				return z
			}
		}else{
			if y<=z{
				return y
			}else {
				return z
			}
		}
	}
	for i:=1;i<n;i++{
		cur1,cur2,cur3:=res[i1]*factors[0],res[i2]*factors[1],res[i3]*factors[2]
		cur:=min(cur1,cur2,cur3)
		if cur==cur1{
			i1++
		}
		if cur==cur2{
			i2++
		}
		if cur==cur3{
			i3++
		}
		res[i]=cur
	}
	return res[n-1]
}
