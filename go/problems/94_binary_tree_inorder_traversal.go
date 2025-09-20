package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	val := []int{}
	left := inorderTraversal(root.Left)
	val = append(val, left...)

	val = append(val, root.Val)

	right := inorderTraversal(root.Right)
	val = append(val, right...)
	return val
}
