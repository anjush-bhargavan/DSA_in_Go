package heap

import "fmt"

type MinHeap struct {
	arr []int
}

func NewHeap() *MinHeap {
	return &MinHeap{arr: []int{}}
}

func (h *MinHeap) Push(val int) {
	h.arr = append(h.arr, val)
	h.ShiftUp(len(h.arr) - 1)
}

func (h *MinHeap) Pop() {
	h.Swap(0, len(h.arr)-1)
	h.arr = h.arr[:len(h.arr)-1]
	h.ShiftDown(0)
}

func (m *MinHeap) Display() {
	fmt.Println(m.arr)
}

func (h *MinHeap) ShiftDown(currIdx int) {
	endIdx := len(h.arr) - 1
	leftIdx := LeftChild(currIdx)
	for leftIdx <= endIdx {
		shiftIdx := leftIdx
		if RightChild(currIdx) <= endIdx && h.arr[leftIdx] > h.arr[RightChild(currIdx)] {
			shiftIdx = RightChild(currIdx)
		}

		if h.arr[shiftIdx] < h.arr[currIdx] {
			h.Swap(shiftIdx, currIdx)
			currIdx = shiftIdx
			leftIdx = LeftChild(currIdx)
		} else {
			break
		}

	}
}

func (h *MinHeap) ShiftUp(currIdx int) {
	parentIdx := ParentIndex(currIdx)

	for currIdx >= 0 && h.arr[currIdx] < h.arr[parentIdx] {
		h.arr[currIdx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[currIdx]
		currIdx = parentIdx
		parentIdx = ParentIndex(currIdx)
	}
}

func (h *MinHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func ParentIndex(i int) int {
	return (i - 1) / 2
}

func LeftChild(i int) int {
	return 2*i + 1
}

func RightChild(i int) int {
	return 2*i + 2
}
