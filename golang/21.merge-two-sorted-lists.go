/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	_h := ListNode {0, nil}
	h := &_h
	
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			h.Next = l1
			l1 = l1.Next
		} else {
			h.Next = l2
			l2 = l2.Next
		}
		h = h.Next
	}

	if l1 != nil {
		h.Next = l1
	} else if l2 != nil {
		h.Next = l2
	}

	return _h.Next
}
// @lc code=end

