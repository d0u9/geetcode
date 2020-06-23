/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func C(left, right int) bool {
	d := left - right
	if d == 0 || d == -1 || d == 1 {
		return true
	}
	return false

}
func recur(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := recur(root.Left)
	if left == -1 {
		return -1
	}
	right := recur(root.Right)
	//fmt.Printf("left = %d, right = %d\n", left, right)
	if right == -1 || !C(left, right){
		return -1
	}
	if (left > right) {
		return left + 1
	}
	return right + 1
}

func isBalanced(root *TreeNode) bool {
	if recur(root) >= 0 {
		return true
	} else {
		return false
	}
}
// @lc code=end

