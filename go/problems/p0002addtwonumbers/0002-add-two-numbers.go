package p0002addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	firstNode := new(ListNode)
	currentNode := firstNode
	next1 := l1
	next2 := l2
	carry := 0
	for next1 != nil || next2 != nil {
		var val1 int
		var val2 int
		if next1 != nil {
			val1 = next1.Val
			next1 = next1.Next
		}
		if next2 != nil {
			val2 = next2.Val
			next2 = next2.Next
		}
		sum := carry + val1 + val2
		currentNode.Next = &ListNode{
			Val: sum % 10,
		}
		currentNode = currentNode.Next
		carry = sum / 10
	}
	if carry > 0 {
		currentNode.Next = &ListNode{
			Val: carry,
		}
	}
	return firstNode.Next
}
