package main

import "fmt"

func main() {
	a := ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	b := removeNthFromEnd(&a, 2)
	fmt.Println(b)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	traveller := head
	lengthList := 0
	for traveller != nil {
		lengthList++
		traveller = traveller.Next
	}
	traveller = head
	var before *ListNode
	after := &ListNode{}
	for i := 1; i <= lengthList-n+1; i++ {
		if i == lengthList-n {
			before = traveller
		}
		if i == lengthList-n+1 {
			after = traveller.Next
		}
		traveller = traveller.Next
	}
	if before == nil {
		return head.Next
	}
	before.Next = after
	return head
}
