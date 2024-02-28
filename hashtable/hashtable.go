package hashtable

import "errors"


type Node struct {
	Key int
	Val int
	Next *Node
}

type Hashtable struct {
	Table []*Node
	Size int
}

func NewHashTable(size int) *Hashtable {
	return &Hashtable{Table: make([]*Node,size),
		Size:size}
}

func (h *Hashtable) hashKey(key int) int {
	return key % h.Size
}

func (h *Hashtable) Put(k,v int) {
	index := h.hashKey(k)

	newNode := &Node{Key: k,Val: v}

	if h.Table[index] != nil {
		newNode.Next = h.Table[index]
		h.Table[index] = newNode
	}else{
		h.Table[index] = newNode
	}
}

func (h *Hashtable) Get(k int) (int,error) {
	index := h.hashKey(k)

	if h.Table[index] != nil {
		curr := h.Table[index]
		for curr != nil {
			if curr.Key == k {
				return curr.Val,nil
			}
			curr = curr.Next
		}
	}
	return 0,errors.New("not found")
}

func (h *Hashtable) Remove(k int) bool {
	index := h.hashKey(k)

	if h.Table[index] != nil {
		var prev *Node
		curr := h.Table[index]
		for curr != nil {
			if curr.Key == k {
				if prev == nil {
					h.Table[index] = curr.Next
				}else{
					prev.Next = curr.Next
				}
				return true
			}
			prev = curr
			curr = curr.Next
		}
	}
	return false
}