package Week02

import "container/list"

//1. barrier node 增加一个标记每层结束的节点
//2. level count  记录每一层的大小
//3. two queues swap 新层用新的队列，旧层没有节点时，新旧队列互换
//4. recursion with pre-order-traverse 前序遍历的方式 获得层次遍历的结果

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder1(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	barrier := &Node{}
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(barrier)
	level := make([]int, 0)
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(*Node)
		if cur == barrier {
			res = append(res, level)
			if queue.Len() > 0 {
				queue.PushBack(barrier)
			}
			level = make([]int, 0)
			continue
		}
		level = append(level, cur.Val)
		for _, n := range cur.Children {
			if n != nil {
				queue.PushBack(n)
			}
		}

	}
	return res
}

func levelOrder2(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		n := queue.Len()
		level := make([]int, 0)
		for i := 0; i < n; i++ {
			cur := queue.Remove(queue.Front()).(*Node)
			level = append(level, cur.Val)
			for _, n := range cur.Children {
				queue.PushBack(n)
			}
		}
		res = append(res, level)
	}
	return res

}

func levelOrder3(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue1 := list.New()
	queue1.PushBack(root)
	queue2 := list.New()
	level := make([]int, 0)
	for queue1.Len() > 0 {
		cur := queue1.Remove(queue1.Front()).(*Node)
		level = append(level, cur.Val)
		for _, n := range cur.Children {
			if n != nil {
				queue2.PushBack(n)
			}
		}
		if queue1.Len() == 0 {
			res = append(res, level)
			queue1, queue2, level = queue2, list.New(), make([]int, 0)
		}

	}
	return res

}

func levelOrder4(root *Node) [][]int {
	res := make([][]int, 0)
	helper(0, root, &res)
	return res
}

func helper(level int, root *Node, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) <= level {
		*res = append(*res, make([]int, 0))
	}
	(*res)[level] = append((*res)[level], root.Val)
	for _, n := range root.Children {
		helper(level+1, n, res)
	}
}
