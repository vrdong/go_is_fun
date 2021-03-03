package main

import "fmt"

func main() {
	mergeKLists([]*ListNode{
		nil,
	})
}

/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
Merge all the linked-lists into one sorted linked-list and return it.
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

Input: lists = []
Output: []

Input: lists = [[]]
Output: []
*/

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

func mergeKLists(lists []*ListNode) *ListNode {
	var result ListNode
	tail := &result
	var min *int
	for {
		min = nil
		current := -1
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				if min == nil {
					min = &lists[i].Val
					current = i
				} else {
					if *min > lists[i].Val {
						min = &lists[i].Val
						current = i
					}
				}
			}
		}
		if current != -1 {
			fmt.Println(lists[current].Val)
			tail.Next = lists[current]
			tail = tail.Next
			lists[current] = lists[current].Next
		} else {
			break
		}
	}
	return result.Next
}
