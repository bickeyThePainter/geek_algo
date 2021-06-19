package Week09

import (
	"container/heap"
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
	left, right := 0, 0
	maxLength := 0
	max := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}
	for _, c := range s {
		if c == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, left*2)
		}
		if left < right {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			right++
		} else {
			left++
		}
		if left == right {
			maxLength = max(maxLength, left*2)
		}
		if left > right {
			left, right = 0, 0
		}
	}
	return maxLength

}


func kmp(a,b string) int{
	if a=="" || b==""{
		return -1
	}
	if len(a)<len(b){
		return -1
	}
	// 计算模式串递归
	var calc func(a string ,pIndex,index int,prefix []int) int
	calc=func(a string,pIndex,index int,prefix []int)int{
		if pIndex==0 && a[pIndex]!=a[index]{
			return 0
		}
		if pIndex==0 && a[pIndex]==a[index]{
			return 1
		}
		if a[pIndex]!=a[index]{
			return calc(a,prefix[pIndex-1],index,prefix)
		}
		return pIndex+1
	}
	n,m:=len(a),len(b)
	prefix:=make([]int,m)
	prefix[0]=0
	// 计算模式串的prefix
	for i:=1;i<m;i++{
		prefix[i]=calc(b,prefix[i-1],i,prefix)
	}
	i,j:=0,0
	for i<n && j<m{
		if a[i]==b[j]{
			i++
			j++
		}else{
			if j==0{
				i++
			}else{
				// 寻找模式串前缀，利用已有成果再匹配 避免重复
				j=prefix[j-1]
			}
		}
	}
	if j==m{
		return i-m
	}
	return -1



}

func findTopKQuery(orders [][]int, k int)  []int {
	durations:=make(map[int]int)
	res:=make([]int,0)
	q:=make(PriorityQueue,0)
	for _,o:=range orders{
		heap.Push(&q,&Item{o[0],o[1]+o[2]})
		durations[o[0]]=o[2]
	}
	for len(res)<k && q.Len()>0{
		cur:=heap.Pop(&q).(*Item)
		res=append(res,cur.queryId)
		heap.Push(&q,&Item{cur.queryId,cur.endTime+durations[cur.queryId]})
		for q.Len()>0 && q[len(q)-1].endTime==cur.endTime{
			_cur:=heap.Pop(&q).(*Item)
			res=append(res,_cur.queryId)
			heap.Push(&q,&Item{
				_cur.queryId,
				_cur.endTime+durations[cur.queryId],
			})
		}
	}
	return res
}

// heap
type Item struct{
	queryId,endTime int
}

type  PriorityQueue []*Item

func (q PriorityQueue) Less(i,j int) bool{
	return q[i].endTime<q[j].endTime
}

func (q PriorityQueue) Swap(i,j int){
	q[i],q[j]=q[j],q[i]
}

func (q PriorityQueue) Len() int{
	return len(q)
}

func (q *PriorityQueue) Pop() interface{}{
	last:=(*q)[len(*q)-1]
	*q=(*q)[0:len(*q)-1]
	return last
}

func (q *PriorityQueue) Push(i interface{}){
	*q=append(*q,i.(*Item))
}



func TestLong(t *testing.T) {
	//longestValidParentheses(
		//")()())")
	//kmp("BBC ABCDAB ABCDABCDABDE","ABCDABD")
	//computeTernary("N?1:Y?4:5")

}
