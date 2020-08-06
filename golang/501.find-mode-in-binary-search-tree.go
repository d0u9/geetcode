/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
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

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func recur(root *TreeNode, mp *map[int]int, max int) int {
	if root == nil {
		return max
	}

	max1 := recur(root.Left,  mp, max)
	max2 := recur(root.Right, mp, max)
	maxc := MaxInt(max1, max2)
	(*mp)[root.Val]++

	return MaxInt(maxc, (*mp)[root.Val])
}

func solution1(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	mp := & map[int]int{}

	max := recur(root, mp, 0)

	for k, v := range(*mp) {
		if v == max {
			ret = append(ret, k)
		}
	}

	return ret
}

func solution2(root *TreeNode) []int {
	ret := []int{}

	

	return ret
}

func findMode(root *TreeNode) []int {
	return solution2(root)
}
// @lc code=end

