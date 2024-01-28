package linkedlist

import "fmt"


type CSNode struct{
	data	int
	next	*CSNode
} 

type CSList struct {
	head *CSNode
}

func (l *CSList) Display(){
	if l.head==nil{
		fmt.Println("empty list")
		return
	}
	curr :=l.head
	for{
		fmt.Println(curr.data)
		curr=curr.next
		if curr==l.head{
			break
		}
	}
}

func (l *CSList) Add(value int){
	newNode :=&CSNode{data: value}

	if l.head==nil{
		l.head=newNode
		newNode.next=l.head
		return
	}
	curr:=l.head
	for curr.next!=l.head{
		curr=curr.next
	}
	curr.next=newNode
	newNode.next=l.head
}

func (l *CSList) Remove(value int) {
	if l.head==nil{
		fmt.Println("nothing to remove, empty list")
		return
	}

	prev:=l.head
	curr:=l.head.next
	for curr.data!= value {
		prev=curr
		curr=curr.next
		if curr==l.head.next{
			fmt.Println("not in the list")
			return
		}
	}

	if curr==l.head{
		if curr.next==l.head{
			l.head=nil
			return
		}else{
			l.head=curr.next
			prev.next=l.head
		}
	}else{
		if curr.next==l.head{
			prev.next=l.head
		}else{
			prev.next=curr.next
		}
	}
}

