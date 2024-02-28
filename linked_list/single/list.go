package single

import "fmt"

//Node holds the model of Node
type Node struct {
	Val  int
	Next *Node
}

//List represents the linkedlist
type List struct {
	Head *Node
	Tail *Node
}

//Add is a method of linkedlist to add data to linked list
func (l *List) Add(data int) {
	newNode := &Node{Val: data}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	l.Tail.Next = newNode
	l.Tail = newNode
}

//Remove is a method of linkedlist to remove a data from linked list
func (l *List) Remove(data int) {
	curr := l.Head
	prev := curr
	for curr != nil && curr.Val != data {
		prev = curr
		curr = curr.Next
	}
	if curr == l.Head {
		l.Head = curr.Next
		return
	}
	prev.Next = curr.Next

}
//Display is a method of linkedlist to display data of linked list
func (l *List) Display(curr *Node) {
	// curr := l.Head

	for curr != nil {
		fmt.Printf("%d > ",curr.Val)
		curr = curr.Next
	}
}

//Middle is a method of linkedlist to find the middle node of linked list
func  Middle(head *Node) *Node {
	if head == nil {
		return nil
	}
	slow := head
	fast := slow.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}


//IsPallindrome is a method of linkedlist to check the nodes of linked list is in palindrome order
func IsPallindrome(head *Node) bool {
	mid := Middle(head)
	second := mid.Next
	mid.Next = nil 

	sec := ReursiveReverse(second)
	
	for head != nil && sec != nil {
		if head.Val != sec.Val {
			return false
		}
		head = head.Next
		sec = sec.Next
	}
	return true
}

//InsertAtIndex is a method of linkedlist to add data at certain index in linked list
func (l *List) InsetAtIndex(data,index int) {
	newNode := &Node{Val: data}
	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		return
	}
	curr := l.Head
	prev := curr
	for i := 0 ; i <= index ; i++ {
		prev = curr
		curr = curr.Next
	}
	prev.Next = newNode
	newNode.Next =curr

}

//Reverse method will reverse the linked list
func (l *List) Reverse() {
	curr := l.Head
	prev := curr
	if curr.Next == nil {
		return
	}
	curr = curr.Next
	prev.Next = nil
	l.Tail = prev
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	l.Head = prev
}


//MergeSort is a method of sorting list nodes
func MergeSort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	
	middle := Middle(head)

	afterMiddle := middle.Next
	middle.Next = nil 

	left := MergeSort(head)
	right := MergeSort(afterMiddle)

	return merge(left,right)

}

func merge(left,right *Node) *Node {
	combined := &Node{}
	curr := combined
	for left != nil && right != nil {
		if left.Val < right.Val {
			curr.Next = left
			left = left.Next
		}else{
			curr.Next = right
			right = right.Next
		}
		curr = curr.Next
	}

	if left != nil {
		curr.Next = left
	}
	if right != nil {
		curr.Next = right
	}
	return combined.Next 
}



func ReursiveReverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReursiveReverse(head.Next)
	
	head.Next.Next = head
	head.Next = nil

	return newHead
}