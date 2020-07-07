/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
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

func match(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}

	if match(a.Left, b.Left) && match(a.Right, b.Right) {
		return true
	}

	return false
}

func recur(s, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if match(s, t) {
		return true
	}
	if recur(s.Left, t) || recur(s.Right, t) {
		return true
	}
	return false
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	return recur(s, t)
}
// @lc code=end

