package tree

import (
	"fmt"
	"math"
)

type Node struct {
	data        int
	left, right *Node
}

type BSTree struct {
	root *Node
}

func (t *BSTree) Insert(value int) bool {
	newNode := &Node{data: value}

	if t.root == nil {
		t.root = newNode
		return true
	}
	temp := t.root
	for {
		if temp.data == value {
			return false
		}
		if temp.data > value {
			if temp.left == nil {
				temp.left = newNode
				return true
			}
			temp = temp.left
		}
		if temp.data < value {
			if temp.right == nil {
				temp.right = newNode
				return true
			}
			temp = temp.right
		}
	}
}

func (t *BSTree) Contains(value int) bool {
	if t.root == nil {
		return false
	}
	temp := t.root
	for temp != nil {
		if temp.data > value {
			temp = temp.left
		} else if temp.data < value {
			temp = temp.right
		} else {
			return true
		}

	}
	return false
}

func (t *BSTree) Remove(value int) {
	t.removeHelper(value, t.root, nil)
}

func (t *BSTree) removeHelper(value int, currentNode, parentNode *Node) {
	for currentNode != nil {
		if value < currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.left
		} else if value > currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.right
		} else {
			if currentNode.right != nil && currentNode.left != nil {
				currentNode.data = GetMinValue(currentNode.right)
				t.removeHelper(currentNode.data, currentNode.right, currentNode)
			} else {
				if parentNode == nil {
					if currentNode.left == nil {
						t.root = currentNode.right
					} else {
						t.root = currentNode.left
					}
				} else {
					if parentNode.right == currentNode {
						if currentNode.right == nil {
							parentNode.right = currentNode.left
						} else {
							parentNode.right = currentNode.right
						}
					} else {
						if currentNode.right == nil {
							parentNode.left = currentNode.left
						} else {
							parentNode.left = currentNode.right
						}
					}
				}
			}
			return
		}
	}
	fmt.Println("not in the tree")
}

func GetMinValue(currentNode *Node) int {
	if currentNode.left == nil {
		return currentNode.data
	} else {
		return GetMinValue(currentNode.left)
	}

}

func (t *BSTree) InOrder() {
	t.inOrderHelper(t.root)
}

func (t *BSTree) inOrderHelper(node *Node) {
	if node != nil {
		t.inOrderHelper(node.left)
		fmt.Printf("%d ", node.data)
		t.inOrderHelper(node.right)
	}
}

func (t *BSTree) PreOrder() {
	t.preOrderHelper(t.root)
}

func (t *BSTree) preOrderHelper(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.data)
		t.preOrderHelper(node.left)
		t.preOrderHelper(node.right)
	}
}

func (t *BSTree) PostOrder() {
	t.postOrderHelper(t.root)
}

func (t *BSTree) postOrderHelper(node *Node) {
	if node != nil {
		t.postOrderHelper(node.left)
		t.postOrderHelper(node.right)
		fmt.Printf("%d ", node.data)
	}
}

func (t *BSTree) FindClosest(target int) int {
	current := t.root
	closest := current.data
	for current != nil {
		if math.Abs(float64(target)-float64(closest)) > math.Abs(float64(target)-float64(current.data)) {
			closest = current.data
		}
		if target < current.data {
			current = current.left
		} else if target > current.data {
			current = current.right
		} else {
			break
		}
	}
	return closest
}

// func(t *BSTree) Validate(){
// 	current := t.root

// 	for current != nil {
// 		current=current.left
// 	}
// 	min := current.data

// 	for current != nil {

// 	}
// }
