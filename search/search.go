package search



func LinearSearch(arr []int,target int) int {
	for i:=0;i<len(arr);i++{
		if arr[i]==target{
			return i
		}
	}
	return -1
}

func BinarySearch(arr []int,target int) int{
	low,high :=0,len(arr)-1

	for low <= high {
		mid:= (low+high)/2

		if arr[mid]==target{
			return mid
		}else if target < arr[mid]{
			high=mid-1
		}else{
			low=mid+1
		}
	}
	return -1
}

func BinarySearchR(arr []int,target,low,high int) int {
	mid:=(low+high)/2
	if low <=high{
		if arr[mid]==target{
			return mid
		}else if target < arr[mid]{
			BinarySearchR(arr,target,low,mid-1)
		}else{
			BinarySearchR(arr,target,mid+1,high)
		}
	}
	return mid
}

func BinarySearchString(s string,target byte) int {
	low,high:=0,len(s)-1
	
	for low <= high{
		mid:=(low+high)/2
		if s[mid]==target{
			return mid
		}else if target < s[mid]{
			high=mid-1
		}else{
			low=mid+1
		}
	} 
	return -1
}