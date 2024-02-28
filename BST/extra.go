package bst

import "math"


func FindDepth(root *Node) int {
	if root == nil {
		return 0
	}
	left := FindDepth(root.Left)
	right := FindDepth(root.Right)

	if left < right {
		return right + 1
	}
	return left + 1
}

func FindHeight(data int,root *Node) int {
	curr := root
	for curr != nil {
		if data < curr.Val {
			curr = curr.Left
		}else if data > curr.Val {
			curr = curr.Right
		}else{
			break
		}
	}
	return FindDepth(curr)
}

var flag = 0

func (t *BST) IsBalanced() bool {
	isBalanceHelper(t.Root)
	return flag == 0
}

func  isBalanceHelper(root *Node) int {
	if root == nil {
		return 0
	}

	left := isBalanceHelper(root.Left)
	right := isBalanceHelper(root.Right)

	if left < right {
		if right - left > 1 {
			flag = 1
		}
		return right + 1
	}
	if left - right > 1 {
		flag = 1
	}
	return left + 1
}

func Validate(node *Node) bool {
	if node == nil {
		return true
	}
	if node.Val <= getMax(node.Left) {
		return false
	} 
	if node.Val >= getMin(node.Right) {
		return false
	}

	if Validate(node.Left) || Validate(node.Right) {
		return false
	}
	return true
}


func getMax(node *Node) int {
	if node.Right == nil {
		return node.Val
	}
	return getMax(node.Right)
}

func FindClosest(node *Node,data int) *Node {
	closest := node 
	curr := node

	for curr != nil {
		if math.Abs(float64(data)-float64(curr.Val)) < math.Abs(float64(data)-float64(closest.Val)) {
			closest = curr
		}
		if data < curr.Val {
			curr  = curr.Left
		}else if data > curr.Val {
			curr = curr.Right
		}else {
			break
		}
	}
	return closest
}


func Successor(data int,root *Node) int{
	result := 0

	curr := root

	for curr != nil {
		if data < curr.Val {
			result = curr.Val
			curr = curr.Left
		}else if data > curr.Val {
			curr = curr.Right
		}else {
			if curr.Right != nil {

				result = getMin(curr.Right)
			}
			break
		}
	}
	return result
}

func SecondMin(root *Node) int {
	secondMin := root.Val
	curr := root
	for curr != nil {
		if curr.Left != nil {
			secondMin = curr.Val
			curr = curr.Left
		}else{
			break
		}
	}
	return secondMin
}