/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head.Next, head
	for fast != nil && slow != nil {
		if fast == slow {
			return true
		}
		if fast.Next == nil {
			return false
		} else {
			fast = fast.Next.Next
		}
		slow = slow.Next
	}
    return false
}
// @lc code=end

