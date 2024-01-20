package linkedlist

// import "fmt"

// type Node struct{
// 	prev *Node
// 	data int
// 	next *Node
// }

// type List struct {
// 	head *Node
// 	tail *Node
// }

// func (l *List) Display(){
// 	if l.head==nil{
// 		fmt.Println("empty")
// 		return
// 	}
// 	curr:=l.head
// 	for curr!=nil{
// 		fmt.Println(curr.data)
// 		curr=curr.next
// 	}
// }

// func (l *List) ReverseDisplay(){
// 	if l.head==nil{
// 		return
// 	}
// 	curr:=l.tail
// 	for curr!=nil{
// 		fmt.Println(curr.data)
// 		curr=curr.prev
// 	}
// }

// func (l *List) Add(value int) {
// 	newNode:=&Node{data: value}
// 	if l.head==nil{
// 		l.head=newNode
// 		l.tail=newNode
// 		return
// 	}
// 	curr:=l.tail
// 	curr.next=newNode
// 	newNode.prev=curr
// 	l.tail=newNode
// }

// func (l *List) Remove(value int){
// 	if l.head== nil{
// 		fmt.Println("nothing to remove")
// 		return
// 	}
// 	if l.head.data==value{
// 		if l.head.next==nil{
// 			l.head=nil
// 			l.tail=nil
// 			return
// 		}
// 		l.head=l.head.next
// 		l.head.prev=nil
// 		return
// 	}
// 	curr:=l.head
// 	for curr.next!=nil && curr.data!=value{
// 		curr=curr.next
// 	}
// 	if curr.next==nil{
// 		if curr.data==value{
// 			l.tail=curr.prev
// 			curr.prev.next=nil
// 			return
// 		}
// 		fmt.Println("the value is not in the list")
// 		return
// 	}
// 	fmt.Printf("data is %d \n",curr.data)
// 	curr.prev.next=curr.next
// 	curr.next.prev=curr.prev

// }
