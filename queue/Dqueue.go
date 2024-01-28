package queue

type DNode struct {
	data int
	next *DNode
}

type Dequeue struct {
	front *DNode
	rear *DNode
}

func (q Dequeue) Display() {

}

func (q Dequeue) PushFront(val int) {

}

func (q Dequeue) PushBack(val int) {

}

func (q Dequeue) PopFront() {

}

func (q Dequeue) PopBack() {
	
}