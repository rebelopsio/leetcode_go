package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var answer []int
	return dfs(root, answer)
}

func dfs(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}
	list = dfs(root.Left, list)
	list = dfs(root.Right, list)
	return append(list, root.Val)
}