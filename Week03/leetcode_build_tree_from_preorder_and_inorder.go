package Week03

// 1. 每次recursion preorder第一个节点为根节点，从inorder中找根节点的坐标，从而确定左子树和右子树的范围。
// 2. 不需要map，记录全局pre指针和in指针，入参stopValue代表树已经生成完毕，每次找到一个stopValue，in++，代表某一个范围的左子树树生成结束，下次该生成右子树。

func buildTree1(preorder []int, inorder []int) *TreeNode {
	check := make(map[int]int)
	for i, k := range inorder {
		check[k] = i
	}
	return build1(preorder, inorder, 0, len(inorder)-1, 0, check)
}

func build1(preorder []int, inorder []int, l, r, m int, check map[int]int) *TreeNode {
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{inorder[l], nil, nil}
	}
	root := &TreeNode{preorder[m], nil, nil}
	sep := check[preorder[m]]
	root.Left = build1(preorder, inorder, l, sep-1, m+1, check)
	root.Right = build1(preorder, inorder, sep+1, r, m+(sep-l)+1, check)
	return root
}

var pre, in int

func buildTree2(preorder []int, inorder []int) *TreeNode {
	pre, in = 0, 0
	return build2(preorder, inorder, -1, false)
}

func build2(preorder []int, inorder []int, stop int, stopValid bool) *TreeNode {
	if in >= len(inorder) || inorder[in] == stop && stopValid {
		in++
		return nil
	}
	root := &TreeNode{preorder[pre], nil, nil}
	pre++
	root.Left = build2(preorder, inorder, root.Val, true)
	root.Right = build2(preorder, inorder, stop, stopValid)
	return root
}
