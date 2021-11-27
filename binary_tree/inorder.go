package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var answer []int
	return dfs(root, answer)

}

func dfs(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}
	list = dfs(root.Left, list)
	list = append(list, root.Val)
	return dfs(root.Right, list)
}