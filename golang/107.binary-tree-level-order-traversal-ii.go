/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
func levelOrderBottom_v1(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	tir := make([]*TreeNode, 0)
	lastlen := 1
	if root == nil {
		return ret
	}
	ret = append(ret, []int{root.Val, })
	tir = append(tir, root)

	for lastlen > 0 {
		lastlen = 0
		cur := make([]int, 0)
		t := make([]*TreeNode, 0)
		for _, e := range tir {
			l, r := e.Left, e.Right
			if l != nil {
				t = append(t, l)
				cur = append(cur, l.Val)
				lastlen++
			}
			if r != nil {
				t = append(t, r)
				cur = append(cur, r.Val)
				lastlen++
			}
		}

		if lastlen > 0 {
			ret = append(ret, cur)
			tir = t
		}
	}

	for i, j := 0, len(ret) - 1; i < len(ret) / 2; i, j = i + 1, j - 1 {
		t := ret[i]
		ret[i] = ret[j]
		ret[j] = t
	}

	return ret
}

func rec(r *TreeNode, ret *[][]int, deep int) {
	if r == nil {
		return
	}
	if len(*ret) <= deep {
		*ret = append(*ret, make([]int, 0))
	}
	a := (*ret)[deep]
	a = append(a, r.Val)
	(*ret)[deep] = a

	rec(r.Left, ret, deep + 1)
	rec(r.Right, ret, deep + 1)
}

func levelOrderBottom_v2(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	rec(root, &ret, 0)
	for i, j := 0, len(ret) - 1; i < len(ret) / 2; i, j = i + 1, j - 1 {
		t := ret[i]
		ret[i] = ret[j]
		ret[j] = t
	}
	return ret
}


func levelOrderBottom(root *TreeNode) [][]int {
	//return levelOrderBottom_v1(root)
	return levelOrderBottom_v2(root)
}
// @lc code=end

