package heap


func BuildHeap(arr []int) {
	for i := (len(arr)/2)-1; i >= 0; i-- {
		HeapifyDown(arr, i)
	}
}

func HeapifyDown(arr []int,i int) {
	for LeftChild(i) <= len(arr) {
		shiftIdx := LeftChild(i)
		if RightChild(i) < len(arr) && arr[shiftIdx] > arr[RightChild(i)] {
			shiftIdx = RightChild(i)
		}
		if arr[shiftIdx] < arr[i] {
			arr[i],arr[shiftIdx] = arr[shiftIdx],arr[i]
		}else{
			return
		}
	}
}

func HeapSort(arr []int) {
	for i := len(arr) ; i >= 0 ; i-- {
		BuildHeap(arr[:i])
		arr[0],arr[i-1] = arr[i-1],arr[0]
	}
}