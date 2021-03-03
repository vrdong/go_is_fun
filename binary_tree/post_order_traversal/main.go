package main

import "fmt"

func main() {
	tree := TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Right: nil,
			Left:  nil,
		},
		Left: &TreeNode{
			Val:   3,
			Right: nil,
			Left:  nil,
		},
	}
	fmt.Println(postorderTraversal(&tree))
}

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	var waiter []TreeNode

	if root != nil {
		waiter = append(waiter, *root)
	}

	for len(waiter) > 0 {
		temp := 0
		current := waiter[len(waiter)-1]
		if current.Right == nil && current.Left == nil {
			result = append(result, current.Val)
			waiter = waiter[:len(waiter)-1]
			continue
		}

		if current.Right != nil {
			waiter[len(waiter)-1].Right = nil
			waiter = append(waiter, *current.Right)
			temp += 1
		}

		if current.Left != nil {
			waiter[len(waiter)-1-temp].Left = nil
			waiter = append(waiter, *current.Left)
		}
	}
	return result
}
