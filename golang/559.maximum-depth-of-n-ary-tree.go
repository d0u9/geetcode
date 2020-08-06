/*
 * @lc app=leetcode id=559 lang=golang
 *
 * [559] Maximum Depth of N-ary Tree
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	maxd := 0
	for _, node := range(root.Children) {
		maxd = MaxInt(maxd, maxDepth(node))
	}
	
	return maxd + 1
}
// @lc code=end

