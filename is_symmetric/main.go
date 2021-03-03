package main

func main() {

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isContrast(root.Left, root.Right)

}

func isContrast(rootLeft *TreeNode, rootRight *TreeNode) bool {
	switch {
	case rootLeft == nil && rootRight == nil:
		return true
	case rootLeft != nil && rootRight != nil:
		return isContrast(rootLeft.Left, rootRight.Right) && isContrast(rootLeft.Right, rootRight.Left)
	default:
		return false
	}
}
