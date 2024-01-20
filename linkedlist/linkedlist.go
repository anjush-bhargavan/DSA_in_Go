package linkedlist

// import "fmt"

// type Node struct{
// 	data int
// 	next *Node
// }

// type List struct{
// 	head *Node
// 	tail *Node
// }

// func(l *List) Display(){
// 	curr:=l.head
// 	if l.head==nil{
// 		fmt.Println("empty list")
// 		return
// 	}
	
// 	for curr!= nil{
// 		fmt.Println(curr.data)
// 		curr=curr.next
// 	}
// }

// func (l *List) Add(value int){
// 	newNode:=&Node{data: value}
// 	if l.head==nil{
// 		l.head=newNode
// 		l.tail=newNode
// 		return
// 	}

// 	l.tail.next=newNode
// 	l.tail=newNode
// 	newNode.next=nil
	
// }

// func (l *List) Remove(value int) {
// 	if l.head==nil {
// 		fmt.Println("there is nothing to remove")
// 		return
// 	}

// 	if l.head.data==value {
// 		if l.head.next==nil{
// 			l.head=nil
// 			l.tail=nil
// 		}
// 		l.head=l.head.next	
// 		return
// 	}

// 	curr :=l.head
// 	for curr.next!=nil && curr.next.data!=value{
// 		curr=curr.next
// 	}
// 	if curr.next==nil{
// 		fmt.Println("the value is not in the list")
// 		return
// 	}
// 	curr.next=curr.next.next
// 	l.tail=curr.next
	
// }

// func (l *List) Reverse(){
// 	if l.head==nil || l.head.next== nil {
// 		return
// 	}
// 	newNode :=l.head
	
// 	curr:=newNode.next
// 	newNode.next=nil
// 	l.tail=newNode
// 	for curr!=nil{
// 		temp:=curr.next
// 		curr.next=newNode
// 		newNode=curr
// 		curr=temp
// 	}
// 	l.head=newNode
// }

// func (l *List) InsertAfter(pos,value int) {
// 	newNode:=&Node{data: value}
// 	if l.head==nil{
// 		fmt.Println("empty list")
// 		return
// 	}
// 	curr:=l.head
// 	for curr.next!=nil && curr.data!=pos{
// 		curr=curr.next
// 	}
// 	if curr.next==nil{
// 		if curr.data==pos{
// 			curr.next=newNode
// 			l.tail=newNode
// 			return
// 		}
// 		fmt.Println("invalid position")
// 		return
// 	}
// 	newNode.next=curr.next
// 	curr.next=newNode
// }

// func (l *List) Search(value int) {
// 	if l.head== nil{
// 		return
// 	}
// 	curr:=l.head
// 	count:=0
// 	for curr.next!=nil && curr.data!=value{
// 		curr=curr.next
// 		count++
// 	}
// 	if curr.next==nil && curr.data!= value{
// 		fmt.Println("the value is not in the list")
// 		return
// 	}
// 	fmt.Printf("\nthe value %d is in position %d\n",value,count)
// }

// func (l *List) RemoveDuplicates(){
// 	if l.head==nil{
// 		return
// 	}
// 	seen:=make(map[int]bool)
// 	prev:=l.head
// 	curr:=l.head.next
// 	for curr!=nil{
// 		if seen[curr.data]{
// 			prev.next=curr.next
// 		}else{
// 			seen[curr.data]=true
// 			prev=curr
// 		}
// 		curr=curr.next
// 	}
// }