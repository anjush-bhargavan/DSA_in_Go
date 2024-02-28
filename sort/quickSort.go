package sort



func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left,right := []int{},[]int{}

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		}else if v > pivot {
			right = append(right,v)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left,pivot),right...)
}