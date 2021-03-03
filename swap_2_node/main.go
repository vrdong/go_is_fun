package main

import "fmt"

func main() {
	swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	})
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	current := head
	isSwap := true
	var tmp int
	for {

		if current == nil {
			break
		}
		fmt.Println(current.Val)
		if isSwap {
			if current.Next != nil {
				tmp = current.Val
				current.Val = current.Next.Val
				current.Next.Val = tmp
			}
			isSwap = false
		} else {
			isSwap = true
		}
		current = current.Next
	}
	return head
}
