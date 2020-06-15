package p0700

import (
	"reflect"
	"testing"
)

func makeRootNode(nums []int) *TreeNode {
	node := new(TreeNode)
	for _, n := range nums {
		node.Insert(n)
	}
	return node
}

func TestSearchBST(t *testing.T) {
	root := makeRootNode([]int{4, 2, 7, 1, 3})
	val := 2
	expected := makeRootNode([]int{2, 1, 3})
	if output := searchBST(root, val); !reflect.DeepEqual(output, expected) {
		t.Fatalf("test failed. output: %#v, expected: %#v", output, expected)
	}

	val2 := 5
	if output := searchBST(root, val2); output != nil {
		t.Fatalf("test failed. output: %#v", output)
	}
}
