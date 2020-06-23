/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minh(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := minh(root.Left), minh(root.Right)
	if left * right  == 0 {
		return max(left, right) + 1
	}
	return min(left, right) + 1
}

func minDepth(root *TreeNode) int {
	return minh(root)
}
// @lc code=end

