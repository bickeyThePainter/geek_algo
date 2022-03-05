package brush

import "testing"

type TreeNode struct {
	    Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
func distributeCoins(root *TreeNode) int {
	res:=0
	dfs(root,&res)
	return res
}

func Abs(x int) int{
	if x<0{
		return -x
	}
	return x
}

func dfs(root *TreeNode,cnt *int) int{
	if root==nil{
		return 0
	}
	l,r:=dfs(root.Left,cnt),dfs(root.Right,cnt)
	*cnt+=Abs(l)
	*cnt+=Abs(r)
	return l+r-1
}


func TestD(t *testing.T){
	root:=&TreeNode{
		Val:   0,
		Left:  &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}
	distributeCoins(root)
}