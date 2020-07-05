/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func curs(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}

	fmt.Println(root)

	left, okl := curs(root.Left, p, q)
	right, okr := curs(root.Right, p, q)

	if left == nil && right == nil {
		if root == p || root == q {
			return root, false
		}
		return nil, false
	}

	if okl == true {
		return left, true
	}
	
	if okr == true {
		return right, true
	}

	if left != nil && right != nil {
		return root, true
	}

	if left == nil && right == nil {
		if root == p || root == q {
			return root, false
		}
		return nil, false
	}

	if left != nil {
		if (left == p && root == q) || (left == q && root == p) {
			return root, true
		}
		return left, false
	}


	if right != nil {
		if (right == p && root == q) || (right == q && root == p) {
			return root, true
		}
		return right, false
	}

	return nil, false
}

// [3,1,4,null,2]\n2\n4
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	r, _ := curs(root, p, q)
	return r
}
// @lc code=end

