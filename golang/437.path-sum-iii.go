/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
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

func recur(root *TreeNode, sum int) (int, []int) {
	if root == nil {
		return 0, []int{}
	}

	c1, arr1 := recur(root.Left, sum)
	c2, arr2 := recur(root.Right, sum)
	cnt := c1 + c2

	arr1 = append(arr1, arr2...)
	for i := 0; i < len(arr1); i++ {
		v := arr1[i] + root.Val
		arr1[i] = v
		if v == sum {
			cnt++
		}
	}
	arr1 = append(arr1, root.Val)
	if root.Val == sum {
		cnt++
	}
	return cnt, arr1
}

func solution1(root *TreeNode, sum int) int {
	r, _ := recur(root, sum)
    return r
}

func solution2(root *TreeNode, sum int) int {
	
}

func pathSum(root *TreeNode, sum int) int {
	return solution1(root, sum)
}
// @lc code=end

