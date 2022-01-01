package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := new(TreeNode)
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 2}
	//root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Println("Max Height of BTree is: ", maxDepth(root))
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	lHeight := maxDepth(root.Left)
	rHeight := maxDepth(root.Right)

	if lHeight >= rHeight {
		return lHeight + 1
	}

	return rHeight + 1
}
