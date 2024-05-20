package bst

import (
	"fmt"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (t *BST) Insert(data int) {
	newNode := &Node{Val: data}
	if t.Root == nil {
		t.Root = newNode
		return
	}
	curr := t.Root
	for {
		if data < curr.Val {
			if curr.Left == nil {
				curr.Left = newNode
				return
			}
			curr = curr.Left
		} else if data > curr.Val {
			if curr.Right == nil {
				curr.Right = newNode
				return
			}
			curr = curr.Right
		} else {
			return
		}
	}
}

func (t *BST) Remove(data int) {
	fmt.Println("Removing node with value :",data)
	removeHelper(data, t.Root)
}

func removeHelper(data int, currNode *Node) *Node {
	if currNode == nil {
		return nil
	}
	if currNode != nil {
		if data < currNode.Val {
			currNode.Left = removeHelper(data, currNode.Left)
		} else if data > currNode.Val {
			currNode.Right =  removeHelper(data, currNode.Right)
		} else {
			if currNode.Left == nil && currNode.Right == nil {
				return nil
			}
			if currNode.Left == nil {
				return currNode.Right
			}else if currNode.Right == nil {
				return currNode.Left
			}
			currNode.Val = GetMin(currNode.Right)
			currNode.Right = removeHelper(currNode.Val,currNode.Right)
		}
	}
	return currNode
}

func GetMin(node *Node) int {
	if node.Left == nil {
		return node.Val
	}
	return GetMin(node.Left)
}

func (t *BST) InOrder() {
	t.inOroderHelper(t.Root)
}

func (t *BST) inOroderHelper(node *Node) {
	if node == nil {
		return
	}
	t.inOroderHelper(node.Left)
	fmt.Printf("%d ", node.Val)
	t.inOroderHelper(node.Right)
}


func (t *BST) Contains(data int) bool {
	curr := t.Root

	for curr != nil {
		if data < curr.Val {
			curr = curr.Left
		}else if data > curr.Val {
			curr = curr.Right
		}else {
			return true
		}
	}
	return false
}



