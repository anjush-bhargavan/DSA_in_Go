package search



func BinarySearch(val int,arr []int) bool {
	low,high := 0,len(arr)

	for low <= high {
		mid := (low + high)/2

		if val < arr[mid] {
			high = mid - 1
		}else if val > arr[mid] {
			low = mid + 1
		}else{
			return true
		}
	} 
	return false
}


func BinarySearchRecursive(arr []int,val,low,high int) int {
	mid := (low+high)/2

	if low <= high {
		if arr[mid] == val {
			return mid
		}else if val < arr[mid] {
			BinarySearchRecursive(arr,val,low,mid-1)
		}else{
			BinarySearchRecursive(arr,val,mid+1,high)
		}
	}
	return -1
}