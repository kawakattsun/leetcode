package p0700

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Insert(val int) {
	if t.Val == 0 {
		t.Val = val
		t.Left = new(TreeNode)
		t.Right = new(TreeNode)
		return
	}
	if t.Val > val {
		t.Left.Insert(val)
	} else {
		t.Right.Insert(val)
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}
