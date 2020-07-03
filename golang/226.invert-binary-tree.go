/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recur(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	recur(root.Left)
	recur(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func invertTree(root *TreeNode) *TreeNode {
    return recur(root)
}
// @lc code=end

