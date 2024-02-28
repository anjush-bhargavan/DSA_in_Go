package queue


type Node struct {
	Val int 
	Next *Node
}

type Queue struct {
	Front *Node
	Rear *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(data int) {
	newNode := &Node{Val: data}
	if q.Front == nil {
		q.Rear = newNode
		q.Front = newNode
		return
	}
	q.Rear.Next = newNode
	q.Rear = newNode 

}

func (q *Queue) Dequeue() *Node {
	node := q.Front
	q.Front = q.Front.Next
	return node
}