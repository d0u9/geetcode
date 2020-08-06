/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
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
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func recur1(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	recur1(root.Left, list)
	*list = append(*list, root.Val)
	recur1(root.Right, list)
}

func solution1(root *TreeNode) int {
	list := make([]int, 0)
	recur1(root, &list)

	// There are at least two nodes in this BST.
	min, l := math.MaxInt32, len(list)
	for i := 1; i < l; i++ {
		min = MinInt(min, (list[i] - list[i - 1]))
	}
	
	return min
}

// return min, max of this subtree
func recur2(root *TreeNode, ret *int) (int, int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}

	min, max := root.Val, root.Val
	d1, d2 := math.MaxInt32, math.MaxInt32
	if root.Left != nil {
		a, b := recur2(root.Left, ret)
		min, max = MinInt(a, min), MaxInt(b, max)
		d1 = root.Val - b
	}

	if root.Right != nil {
		a, b := recur2(root.Right, ret)
		min, max = MinInt(a, min), MaxInt(b, max)
		d2 = a - root.Val
	}

	*ret = MinInt((*ret), MinInt(d1, d2))

	return min, max
}

func solution2(root *TreeNode) int {
	ret := math.MaxInt32
	recur2(root, &ret)
	return ret
}

func getMinimumDifference(root *TreeNode) int {
    return solution2(root)
}
// @lc code=end

