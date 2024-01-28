package linkedlist

import "fmt"

type DNode struct{
	prev *DNode
	data int
	next *DNode
}

type DList struct {
	head *DNode
	tail *DNode
}

func (l *DList) Display(){
	if l.head==nil{
		fmt.Println("empty")
		return
	}
	curr:=l.head
	for curr!=nil{
		fmt.Println(curr.data)
		curr=curr.next
	}
}

func (l *DList) ReverseDisplay(){
	if l.head==nil{
		return
	}
	curr:=l.tail
	for curr!=nil{
		fmt.Println(curr.data)
		curr=curr.prev
	}
}

func (l *DList) Add(value int) {
	newDNode:=&DNode{data: value}
	if l.head==nil{
		l.head=newDNode
		l.tail=newDNode
		return
	}
	curr:=l.tail
	curr.next=newDNode
	newDNode.prev=curr
	l.tail=newDNode
}

func (l *DList) Remove(value int){
	if l.head== nil{
		fmt.Println("nothing to remove")
		return
	}
	if l.head.data==value{
		if l.head.next==nil{
			l.head=nil
			l.tail=nil
			return
		}
		l.head=l.head.next
		l.head.prev=nil
		return
	}
	curr:=l.head
	for curr.next!=nil && curr.data!=value{
		curr=curr.next
	}
	if curr.next==nil{
		if curr.data==value{
			l.tail=curr.prev
			curr.prev.next=nil
			return
		}
		fmt.Println("the value is not in the Dlist")
		return
	}
	fmt.Printf("data is %d \n",curr.data)
	curr.prev.next=curr.next
	curr.next.prev=curr.prev

}
