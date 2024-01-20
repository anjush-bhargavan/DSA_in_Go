package heap

import (
	"fmt"
)

type Heap struct{
	arr []int
}


// func (h *Heap) BuildHeap(array []int) {
// 		heap := &Heap{arr: array}
// 	for i := len(heap.arr)/2 - 1 ; i >= 0 ; i-- {
// 		heapify(array, i , len(heap.arr))
// 	}
// }

// func heapify (array []int, i,n int) {
// 	largest := i
// 	left := 2 * i + 1
// 	right := 2 * i + 2

// 	if left < n && array[left] < array[largest] {
// 		largest = left
// 	}

// 	if right < n && array[right] < array[largest] {
// 		largest = right
// 	}

// 	if largest != i {
// 		array[i],array[largest] = array[largest],array[i]
// 		heapify(array, largest,n)
// 	}
// }

func (h *Heap) BuildHeap(array []int) {
	h.arr = array

	for i := h.parent(len(h.arr)-1); i >= 0 ; i-- {
		h.shiftDown(i)
	}
}

func (h *Heap) shiftDown(currentIdx int) {
	endIdx := len(h.arr)-1
	leftIdx := h.leftChild(currentIdx)
	for leftIdx <= endIdx {
		rightIdx := h.rightChild(currentIdx)
		idxToShift := leftIdx
		if rightIdx <= endIdx && h.arr[rightIdx] > h.arr[leftIdx] {
			idxToShift = rightIdx
		}
		
		if h.arr[currentIdx] < h.arr[idxToShift] {
			h.arr[currentIdx],h.arr[idxToShift] = h.arr[idxToShift],h.arr[currentIdx]
			currentIdx = idxToShift
			leftIdx = h.leftChild(currentIdx)
		}else {
			return
		}
	}
}

func (h *Heap) shiftUp(currentIdx int) {
	parentIdx := h.parent(currentIdx)
	for currentIdx > 0 && h.arr[parentIdx] < h.arr[currentIdx] {
		h.arr[currentIdx],h.arr[parentIdx] = h.arr[parentIdx],h.arr[currentIdx]
		currentIdx = parentIdx
		parentIdx = h.parent(currentIdx)
	}
}

func (h *Heap) Peek() int {
	return h.arr[0]
}

func (h *Heap) Remove() {
	h.arr[0],h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[0]
	h.arr = h.arr[:len(h.arr)-1]
	h.shiftDown(0)
	
}

func (h *Heap) Insert(value int) {
	h.arr = append(h.arr, value)
	h.shiftUp(len(h.arr)-1)
}

func (h *Heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) leftChild(i int) int {
	return 2 * i + 1
}

func (h *Heap) rightChild(i int) int {
	return 2 * i + 2
}

func (h *Heap) Display() {
	fmt.Println(h.arr)
}


func(h *Heap) HeapSort(arr []int) {
	 h.BuildHeap(arr)
	for i := 0 ; i<len(arr)-i ; i ++ {
		arr[0],arr[len(arr)-1-i] = 	arr[len(arr)-1-i],arr[0]
		 
	}
}