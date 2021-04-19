package Week03

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. 递归 由于p,q一定存在在tree中，所以root和p,q其中一个相等就可以返回，不需继续递归左右子树，否则递归左右子树，如果都不为空返回root 否则返回不为空的一个
// 2. 父节点路径 记录p,q的父节点路径 看是否有交点

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else {
		return right
	}
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	// parent map
	if root == nil || root == p || root == q {
		return root
	}
	check := make(map[*TreeNode]*TreeNode)
	stack := list.New()
	stack.PushFront(root)
	for check[p] == nil || check[q] == nil {
		cur := stack.Remove(stack.Front()).(*TreeNode)
		if cur.Left != nil {
			check[cur.Left] = cur
			stack.PushFront(cur.Left)
		}
		if cur.Right != nil {
			check[cur.Right] = cur
			stack.PushFront(cur.Right)
		}
	}
	// 以下 同判断两个链表的交点
	p1 := p
	q1 := q
	for p != q {
		if p == nil {
			p = q1
		} else {
			p = check[p]
		}
		if q == nil {
			q = p1
		} else {
			q = check[q]
		}
	}
	return q
	/*set := make(map[*TreeNode]struct{})
	for p != nil {
		set[p] = struct{}{}
		p = check[p]
	}
	for q != nil {
		if _, ok := set[q]; ok {
			return q
		}
		q = check[q]
	}
	return q*/
}
