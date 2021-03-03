/* Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
According to the definition of LCA on WIKIPEDIA. The lowest common ancestor is defined between 2 nodes p and q as the
lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself)
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// If no more node
	_, result := find(root, p, q)
	return result
}

func find(root, p, q *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	findLeft, lcaLeft := find(root.Left, p, q)
	findRight, lcaRight := find(root.Right, p, q)

	if lcaLeft != nil {
		return 2, lcaLeft
	}

	if lcaRight != nil {
		return 2, lcaRight
	}

	if findLeft+findRight == 2 {
		return 2, root
	}

	if root.Val == q.Val || root.Val == p.Val {
		if findLeft+findRight == 1 {
			return 2, root
		} else {
			return 1, nil
		}
	}

	return findLeft + findRight, nil
}

func main() {

}
