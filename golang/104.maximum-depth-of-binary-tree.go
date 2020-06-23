/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return maxInt(recur(root.Left), recur(root.Right)) + 1
}

func maxDepth(root *TreeNode) int {
	return recur(root)
}
// @lc code=end

