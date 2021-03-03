package main

import "fmt"

func main() {
	reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}, 3)
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	traveler := head
	count := 0
	for traveler != nil {
		if count%k == 0 {
			swapValueInKNode(traveler, k)
		}
		count += 1
		traveler = traveler.Next
	}
	return head
}

func swapValueInKNode(head *ListNode, k int) {
	fmt.Println(head.Val)
	if k == 1 {
		return
	}
	if head == nil {
		return
	}
	count := 1
	traveler := head
	for traveler != nil {
		traveler = traveler.Next
		if traveler != nil {
			count += 1
		}
		if count == k {
			break
		}
	}
	if traveler == nil {
		return
	}

	tmp := head.Val
	head.Val = traveler.Val
	traveler.Val = tmp
	swapValueInKNode(head.Next, k-2)
}
