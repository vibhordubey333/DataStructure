package main

import (
	"fmt"
)

/*
	    1
	   / \
	  2   3
	 /   / \
	4   5	6
	Tree Height: 3
*/

type TreeNode struct {
	value     interface{}
	leftNode  *TreeNode
	rightNode *TreeNode
}

func (t *TreeNode) initializeTreeNode() *TreeNode {
	root := new(TreeNode)
	root.value = 1
	root.leftNode = &TreeNode{value: 2}
	root.rightNode = &TreeNode{value: 3}

	root.leftNode.leftNode = &TreeNode{value: 4}
	root.rightNode.leftNode = &TreeNode{value: 5}
	root.rightNode.rightNode = &TreeNode{value: 6}
	return root
}

func (t *TreeNode) findMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.leftNode == nil && root.rightNode == nil {
		return 1
	}
	leftHeight := root.findMaxDepth(root.leftNode)
	rightHeight := root.findMaxDepth(root.rightNode)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func main() {
	treeNodeObject := new(TreeNode)
	treeNodeObject = treeNodeObject.initializeTreeNode()
	treeHeight := treeNodeObject.findMaxDepth(treeNodeObject)
	fmt.Println("Maximum Height Of A Binary Tree: ", treeHeight)
}
