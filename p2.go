package leetcode_go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return recurAddTwoNumbers(l1, l2, 0)
}

func recurAddTwoNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry > 0 {
			return &ListNode{
				Next: nil,
				Val: carry,
			}
		}
		return nil
	}

	l1val := 0
	var l1next *ListNode = nil

	if l1 != nil {
		l1val = l1.Val
		l1next = l1.Next
	}

	l2val := 0
	var l2next *ListNode = nil

	if l2 != nil {
		l2val = l2.Val
		l2next = l2.Next
	}

	sum := l1val + l2val + carry
	carry = 0

	if sum >= 10 {
		carry = sum / 10
		sum -= 10
	}

	return &ListNode{
		Val: sum,
		Next: recurAddTwoNumbers(l1next, l2next, carry),
	}
}