package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	result := inorderTraversal(&tree)
	fmt.Println(result)
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var waiter []TreeNode

	if root != nil {
		waiter = append(waiter, *root)
	}

	for len(waiter) > 0 {

		current := waiter[len(waiter)-1]

		if current.Left != nil {
			waiter[len(waiter)-1].Left = nil
			waiter = append(waiter, *current.Left)
			continue
		}

		result = append(result, current.Val)
		waiter = waiter[:len(waiter)-1]
		if current.Right != nil {
			waiter = append(waiter, *current.Right)
		}
	}
	return result
}
