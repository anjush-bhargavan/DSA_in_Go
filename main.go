package main

import (
	// "dsa/sort"

	// "dsa/recursion"
	// tree "dsa/tree/bst"
	// "dsa/search"
	// "dsa/stack"
	// "dsa/heap"
	"dsa/graph"
	// "dsa/trie"
	"fmt"
)

// "dsa/linkedlist"
// "dsa/search"
// "fmt"



func main(){
	
	g := &graph.Graph{
		Vertex: make(map[int][]int),
	} 
	g.Insert(3,4,false)
	g.Insert(3,2,false)
	// g.Insert(4,1,false)
	g.Insert(2,4,false)
	g.Insert(2,1,false)
	g.Insert(1,5,false)
	// g.RemoveVertex(2)
	// g.RemoveEdge(2,4,true)
	// g.DFS(2)
	fmt.Println(g.IsPathExists(4,2))
	// g.Display()
 
	// trie := trie.NewSuffixTrie()

	// trie.Trie("black")
	// trie.Trie("big bang")
	// trie.PrefixString("paper")
	// trie.Trie("amazon")
	// trie.Trie("amazing")

	// fmt.Println(trie.Contains("pap"))

	// arr:=[]int{3,7,6,15,8,11,9,44,35,2,1}

	// h :=&heap.Heap{}

	// h.BuildHeap(arr)
	// h.Insert(10)
	// h.Insert(15)
	// h.Insert(1)
	// h.Insert(2)
	// h.Remove()
	// h.BuildHeap(arr)
	// fmt.Println(h.Peek())
	// h.Display()
	// fmt.Println(arr)


	// sort.SelectionSort(arr)
	// sort.BubbleSort(arr)
	// sort.InsertionSort(arr)
	// result:=sort.MergeSort(arr)
	// sort.QuickSort(arr)
	// fmt.Println(arr)

	// tree := &tree.BSTree{}
	// fmt.Println()
	// tree.Insert(10)
	// tree.Insert(5)
	// tree.Insert(16)
	// tree.Insert(14)
	// tree.Insert(3)
	// tree.Insert(6)
	// tree.Insert(12)
	// tree.Insert(8)
	// tree.Insert(8)
	// tree.Insert(18)
	// tree.Remove(10)
	// if tree.Contains(10){
	// 	fmt.Println("in")
	// }else{
	// 	fmt.Println("no")
	// }
	// tree.InOrder()
	// fmt.Println("")
	// tree.PreOrder()
	// fmt.Println("")
	// tree.PostOrder()
	// fmt.Println(tree.FindClosest(13))


	// t:=recursion.FibonacciSeries(8)
	// fmt.Println(t)
	// index:=search.BinarySearchString("abcdef",'d')
	// fmt.Println(index)

	// arr :=[]int{1,2,3,4,5,6,7,8}
	// fmt.Println(search.BinarySearchR(arr,5,0,7))

	// list:=&linkedlist.List{} 

	// // list.Add(8)
	// list.Add(5)
	// list.Add(2)
	// list.Add(4)
	// list.Add(1)
	// list.Add(3)
	// list.Add(1)
	// list.Add(1)
	// list.Remove(3)
	// list.InsertAfter(1,7)
	// list.Reverse()
	// list.Search(1)
	// list.RemoveDuplicates()
	// list.Display()
	// list.ReverseDisplay()

	// stack := &stack.Stack{}
	// stack.Push(1)
	// stack.Push(2)
//  stack.Push(3)
// 	stack.Pop()
// 	stack.Pop()
// data:=stack.Peak()
// fmt.Println(data)
// if stack.Isempty(){
// 	fmt.Println("empty")
// }else{
// 	fmt.Println("not empty")
// }
// 	stack.DisplayStack()

	// queue :=&queue.Queue{}
	// queue.Enqueue(1)
	// queue.Enqueue(2)
	// queue.Enqueue(3)
	// queue.Enqueue(4)

	// queue.Dequeue()
	// queue.Dequeue()
	// queue.DisplayQueue()

	// s:= string.Encode("hai",28)
	// fmt.Println(s)

	// n:=recursion.Factorial(5)
	// fmt.Println(n)

	// recursion.Sample(5)
}