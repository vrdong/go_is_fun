package main

import "fmt"

func main() {
	a := mergeTwoLists(&ListNode{
		Val:  0,
		Next: nil,
	}, &ListNode{
		Val:  0,
		Next: nil,
	})
	fmt.Println(a)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, traveller *ListNode
	result = &ListNode{}
	traveller = result
	for l1 != nil || l2 != nil {
		if l1 == nil {
			traveller.Next = l2
			break
		}
		if l2 == nil {
			traveller.Next = l1
			break
		}
		if l1.Val > l2.Val {
			traveller.Next = l2
			l2 = l2.Next
		} else {
			traveller.Next = l1
			l1 = l1.Next
		}
		traveller = traveller.Next
	}
	return result.Next
}
