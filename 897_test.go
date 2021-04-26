// 给你一棵二叉搜索树，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var startNode = &TreeNode{}
var currentNode = startNode

// todo: 重写
func increasingBST(root *TreeNode) *TreeNode {
	midOrderTraversal(root)
	return startNode.Right
}

func midOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}
	midOrderTraversal(node.Left)
	currentNode.Left = node
	node.Left = nil
	currentNode = node.Right
	midOrderTraversal(node.Right)
}
