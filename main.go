package main

import (
	// "fmt"

	// "github.com/anjush-bhargavan/dsa/trie"
	// "github.com/anjush-bhargavan/dsa/search"
	// "github.com/anjush-bhargavan/dsa/hashtable"
	// "github.com/anjush-bhargavan/dsa/sort"
	// "github.com/anjush-bhargavan/dsa/queue"
	bst "github.com/anjush-bhargavan/dsa/BST"
	// "fmt"
	// "github.com/anjush-bhargavan/dsa/graph"
	// "github.com/anjush-bhargavan/dsa/heap"
	// "github.com/anjush-bhargavan/dsa/linked_list/single"
)
func main() {
	arr := []int{6, 3, 2, 17, 1, 9, 78, 6}
	// heap.HeapSort(arr)
	// fmt.Println(arr)
	// result := sort.MergeSort(arr)
	// result := sort.QuickSort(arr)
	// sort.BubbleSort(arr)
	// sort.InsertionSort(arr)
	// sort.SelectionSort(arr)
	// fmt.Println(result)
	// fmt.Println(search.BinarySearchRecursive(arr,16,0,len(arr)))
	// q := queue.NewDeQueue()
	// for _, v := range arr {
	// 	q.PushRear(v)
	// }
	// // q.PopFront()
	// q.PushFront(4)
	// queue.Display(q.Front)

	// h := hashtable.NewHashTable(5)

	// h.Put(1,11)
	// h.Put(2,22)
	// h.Put(3,33)
	// h.Put(6,66)
	// h.Put(7,77)
	// h.Remove(7)

	// fmt.Println(h.Get(7))






	// l := &single.List{}
	// for _, v := range arr {
		
	// 	l.Add(v)
	// }
	// fmt.Println("hi")
	// // l.InsetAtIndex(10,2)
	// // l.Reverse()
	// rev := single.MergeSort(l.Head)
	// fmt.Println("middle is :",single.Middle(rev).Val)
	// new := single.ReursiveReverse(l.Head)
	// l.Display(new)
	// fmt.Println(single.IsPallindrome(l.Head))

	t := &bst.BST{}

	// m := heap.NewHeap()
	for _,v := range arr {
	// m.Push(v)
	t.Insert(v)
	}
	t.Remove(3)
	t.InOrder()
	// x := 65
	// fmt.Println(t.IsBalanced())
	// fmt.Println("\nHeight of the tree",bst.FindDepth(t.Root))
	// fmt.Printf("\nDepth of the node %d is %d\n",7,bst.FindHeight(9,t.Root))
	// fmt.Printf("\nCloset node to %d is %d",4,bst.FindClosest(t.Root,11).Val)
	// fmt.Printf("\n the successor of %d is %d ",x,bst.Successor(x,t.Root))
	// fmt.Println("Second min is : ",bst.SecondMin(t.Root))
	// m.Pop()
	// m.Pop()
	// m.Display()
	// fmt.Println(m.)

	// g := graph.NewGraph()

	// g.Insert(1, 2)
	// g.Insert(3, 4)
	// g.Insert(1, 3)
	// g.Insert(1, 5)
	// g.Insert(1, 7)
	// g.Insert(2, 4)
	// g.Insert(8, 1)

	// g.BFS(8)
	// visited := make(map[int]bool)
	// fmt.Println(g.IsPathExists(8, 7))
	// g.DFS(8,visited)
	// g.RemoveEdge(1,5)
	// g.RemoveVertex(1)
	// g.Display()

	// t := trie.NewTrie()

	// t.Insert("bigbang")
	// t.Insert("big")

	// t.Remove("bigbang")
	// t.SuffixTrie("bigbang")
	// fmt.Println(t.Contains("hinadh"))
	// fmt.Println(t.Contains("hello"))
	// fmt.Println(t.Contains("bigbang"))
	// fmt.Println(t.Contains("big"))

}
