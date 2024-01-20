package stack

import "fmt"

type Node struct{
	data int
	next *Node
}

type Stack struct{
	top *Node

}

func(s *Stack) DisplayStack() {
	if s.top==nil{
		fmt.Println("stack underflow")
		return
	}
	curr := s.top
	for curr!=nil{
		fmt.Println(curr.data)
		curr=curr.next
	}
}

func (s *Stack) Push(value int) {
	newNode :=&Node{data: value}
	if s.top==nil{
		s.top=newNode
		return
	}
	newNode.next=s.top
	s.top=newNode
}

func(s *Stack) Pop(){
	if s.top==nil{
		fmt.Println("stack underflow")
		return
	}
	s.top=s.top.next
}

func(s *Stack) Peak() int{
	if s.top==nil{
		return -1
	}
	return s.top.data
}

func (s *Stack) Isempty()bool{
	return s.top==nil
}

// func Length(s Stack) int{
// 	return l
// }