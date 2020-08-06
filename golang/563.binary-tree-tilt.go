/*
 * @lc app=leetcode id=563 lang=golang
 *
 * [563] Binary Tree Tilt
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
func Abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

func recur(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	ret1, sum1 := recur(root.Left)
	ret2, sum2 := recur(root.Right)

	ret := ret1 + ret2 + Abs(sum1 - sum2)
	sum := sum1 + sum2 + root.Val

	return ret, sum
}

func findTilt(root *TreeNode) int {
   v, _ := recur(root)
   return v
}
// @lc code=end

