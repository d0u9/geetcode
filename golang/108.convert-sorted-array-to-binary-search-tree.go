/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func tree(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}

	m := l / 2
	left := tree(nums[:m])
	right := tree(nums[m+1:])
	return &TreeNode {
		Val: nums[m],
		Left: left,
		Right: right,
	}
}

func tree2(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	m := (left + right) / 2
	return &TreeNode {
		Val: nums[m],
		Left: tree2(nums, left, m - 1),
		Right: tree2(nums, m + 1, right),
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	//return tree(nums)
	
	return tree2(nums, 0, len(nums) - 1)
}
// @lc code=end

