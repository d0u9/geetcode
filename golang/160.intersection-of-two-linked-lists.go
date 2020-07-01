/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func solution1(headA, headB *ListNode) *ListNode {
	l1, l2 := 0, 0
	cur1, cur2 := headA, headB
	for cur1 != nil {
		l1++
		cur1 = cur1.Next
	}
	for cur2 != nil {
		l2++
		cur2 = cur2.Next
	}
	// l1 - long list, l2 - short list
	if l2 > l1 {
		l1, l2 = l2, l1
		cur1, cur2 = headB, headA
	} else {
		cur1, cur2 = headA, headB
	}
	for i := 0; i < l1 - l2; i++ {
		cur1 = cur1.Next
	}
	for cur1 != nil && cur2 != nil {
		if cur1 == cur2 {
			return cur1
		}
		cur1, cur2 = cur1.Next, cur2.Next 
	}
	return nil
}

func solution2(headA, headB *ListNode) *ListNode {
	cur1, cur2 := headA, headB
	for cur1.Next != nil && cur2.Next != nil {
		if cur1 == cur2 {
			return cur1
		}
		cur1, cur2 = cur1.Next, cur2.Next
	}
	if cur1 == nil {
		cur1.Next = headA
	} else {
		cur2.Next = headB
	}
	for cur1 != nil && cur2 != nil {
		if cur1 == cur2 {
			return cur1
		}
		if cur1.Next == nil {
			return nil
		}
		cur1 = cur1.Next.Next
		cur2 = cur2.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	return solution1(headA, headB)
}
// @lc code=end
