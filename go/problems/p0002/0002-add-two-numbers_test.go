package p0002

import (
	"reflect"
	"testing"
)

func makeListNode(nums []int) *ListNode {
	node := new(ListNode)
	currentNode := node
	for _, n := range nums {
		currentNode = &ListNode{Val: n}
		currentNode = currentNode.Next
	}
	return node.Next
}

func TestAddTwoNumbers(t *testing.T) {
	cases := [][]*ListNode{
		{
			makeListNode([]int{2, 4, 3}),
			makeListNode([]int{5, 6, 4}),
			makeListNode([]int{7, 0, 8}),
		},
		{
			makeListNode([]int{1, 8}),
			makeListNode([]int{0}),
			makeListNode([]int{1, 8}),
		},
	}

	for _, c := range cases {
		l1 := c[0]
		l2 := c[1]
		expected := c[2]
		if output := addTwoNumbers(l1, l2); !reflect.DeepEqual(output, expected) {
			t.Fatalf("test failed. output: %#v, expected: %#v", output, expected)
		}
	}
}
