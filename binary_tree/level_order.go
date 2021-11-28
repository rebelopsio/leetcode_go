package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    var q []*TreeNode
    q = append(q, root)

    for len(q) > 0 {
        length := len(q)
        level := []int{}

        for i := 0; i < length; i++ {
            node := q[0]
            q = q[1:]

            level = append(level, node.Val)
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        if len(level) > 0 {
            result = append(result, level)
        }
    }
    return result
}