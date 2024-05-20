package bst

import "math"


func FindHeight(Root *Node) int {
	if Root == nil {
		return -1
	}
	left := FindHeight(Root.Left)
	right := FindHeight(Root.Right)

	return max(left,right) + 1
}


func FindDepth(Root *Node,val int,depth int) int {
	if Root == nil {
		return -1
	}
	if Root.Val == val {
		return depth
	}
	leftDepth := FindDepth(Root.Left,val,depth+1)
	if leftDepth != -1 {
		return leftDepth
	}
	
	return FindDepth(Root.Right,val,depth+1)
} 


func IsBalanced(root *Node) bool {
	left := FindHeight(root.Left)
	right := FindHeight(root.Right)

	return max(left,right) - min(left,right) <= 1
}

func IsValid(root *Node) bool {
	return root.Val < GetMin(root.Right) && root.Val > GetMax(root.Left) 
}


func GetMax(node *Node) int {
	if node.Left == nil {
		return node.Val
	}
	return GetMax(node.Left)
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

				result = GetMin(curr.Right)
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