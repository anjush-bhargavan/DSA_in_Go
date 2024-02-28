package queue

import "fmt"

type Dequeue struct {
	Front *Node
	Rear  *Node
}

func NewDeQueue() *Dequeue {
	return &Dequeue{}
}

func (d *Dequeue) PushFront(data int) {
	newNode := &Node{Val: data}
	if d.Front == nil {
		d.Front = newNode
		d.Rear = newNode
		return
	}
	newNode.Next = d.Front
	d.Front = newNode
}

func (d *Dequeue) PushRear(data int) {
	newNode := &Node{Val: data}
	if d.Front == nil {
		d.Front = newNode
		d.Rear = newNode
		return
	}
	d.Rear.Next = newNode
	d.Rear = newNode
}

func (d *Dequeue) PopFront() *Node {
	node := d.Front
	d.Front = d.Front.Next
	return node
}

func (d *Dequeue) PopRear() *Node {
	node := d.Rear
	d.Rear = nil
	curr := d.Front
	for curr != node {
		d.Rear = curr
		curr = curr.Next
	}
	d.Rear.Next = nil
	return node
}

func Display(front *Node) {
	curr := front
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
}
