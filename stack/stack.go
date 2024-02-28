package stack


type Stack struct {
	Top *Node
}


type Node struct {
	Val interface{}
	Next *Node
}


func (s *Stack) Push(data any) {
	newNode := &Node{Val: data}
	newNode.Next = s.Top
	s.Top = newNode
}

func (s *Stack) Pop() {
	s.Top = s.Top.Next
}

func (s *Stack) Peek() any {
	return s.Top.Val
}