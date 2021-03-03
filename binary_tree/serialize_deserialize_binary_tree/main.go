//Serialize and Deserialize Binary Tree
/* Design an algorithm to serialize and deserialize a binary tree*/

package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
// serialize we use formula root*left*right
// 4*-7*_*_*-3*_*_
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "_"
	}
	return fmt.Sprintf("%d", root.Val) + "*" + this.serialize(root.Left) + "*" + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
//deserialize we use recursion to deserialize base on concept leaf = root + 1
func (this *Codec) deserialize(data string) *TreeNode {
	fmt.Println(data)
	if data[0] == '_' {
		return nil
	}
	rootTail := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '*' {
			rootTail = i
			break
		}
	}
	val, err := strconv.Atoi(data[0:rootTail])
	if err != nil {
		return nil
	}
	countRoot := 0
	countLeaf := 0
	index := 0
	for i := rootTail + 1; i < len(data); i++ {
		if data[i] == '*' {
			if data[i-1] == '_' {
				countLeaf += 1
			} else {
				countRoot += 1
			}
		}
		if countRoot == countLeaf-1 {
			index = i
			break
		}
	}

	root := &TreeNode{
		Val:   val,
		Left:  this.deserialize(data[rootTail+1 : index+1]),
		Right: this.deserialize(data[index+1:]),
	}
	return root
}

func main() {
	ser := Constructor()
	deSer := Constructor()
	data := ser.serialize(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	})
	fmt.Println(data)
	ans := deSer.deserialize("4*-7*_*_*-3*_*_")
	fmt.Println(ans)
}
