package main

import "fmt"

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func main() {
	tree := TreeNode{
		Val: 3,
		Right: &TreeNode{
			Val: 20,
			Right: &TreeNode{
				Val:   7,
				Right: nil,
				Left:  nil,
			},
			Left: &TreeNode{
				Val:   15,
				Right: nil,
				Left:  nil,
			},
		},
		Left: &TreeNode{
			Val:   9,
			Right: nil,
			Left:  nil,
		},
	}

	fmt.Println(levelOrder(&tree))
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var queue []TreeNode

	var level []int
	front := -1
	if root != nil {
		queue = append(queue, *root)
		level = append(level, 0)
		front += 1
	} else {
		return result
	}

	rear := 0
	for front <= rear {
		current := queue[front]
		currentLevel := level[front]
		if currentLevel > len(result)-1 {
			result = append(result, []int{})
		}
		result[currentLevel] = append(result[currentLevel], current.Val)

		if current.Left != nil {
			queue = append(queue, *current.Left)
			level = append(level, currentLevel+1)
			rear += 1
		}

		if current.Right != nil {
			queue = append(queue, *current.Right)
			level = append(level, currentLevel+1)
			rear += 1
		}
		front += 1
	}
	return result
}
