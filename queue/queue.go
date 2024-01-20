package queue

import "fmt"

type Node struct{
	data int
	next *Node
}

type Queue struct{
	front *Node
	rear *Node
}

func (q *Queue) Enqueue(value int) {
	newNode :=&Node{data: value}
	if q.front== nil{
		q.front=newNode
		q.rear=newNode
	}

	q.rear.next=newNode
	q.rear=newNode
}

func (q *Queue) Dequeue(){
	if q.front==nil{
		fmt.Println("empty queue")
		return
	}

	q.front=q.front.next
}

func (q *Queue) DisplayQueue(){
	if q.front==nil {
		fmt.Println("empty queue")
		return
	}

	curr := q.front
	for curr!=nil{
		fmt.Println(curr.data)
		curr=curr.next
	}
}