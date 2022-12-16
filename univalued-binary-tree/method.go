package univaluedbinarytree

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	var leftCheck, rightCheck bool
	if root.Left == nil && root.Right == nil {
		return true
	}

	//check left
	if root.Left == nil {
		leftCheck = true
	} else {
		if root.Left.Val != root.Val {
			return false
		}
		leftCheck = isUnivalTree(root.Left)
	}

	//check right
	if root.Right == nil {
		rightCheck = true
	} else {
		if root.Right.Val != root.Val {
			return false
		}
		rightCheck = isUnivalTree(root.Right)
	}

	if leftCheck && rightCheck {
		return true
	}

	return false
}
