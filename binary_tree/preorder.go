package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
	var answer []int
	return dfs(root, answer)

}

func dfs(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}
	list = append(list, root.Val)
	list = dfs(root.Left, list)
	return dfs(root.Right, list)
}