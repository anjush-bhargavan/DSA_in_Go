package sort


func Pivot(arr []int, pivotIndex,endIndex int) int {
	swapIndex:=pivotIndex
	for i:=pivotIndex+1;i<=endIndex;i++{
		if arr[i] < arr[pivotIndex]{
			swapIndex++
			arr[swapIndex],arr[i]=arr[i],arr[swapIndex]
		}
		
	}
	arr[pivotIndex],arr[swapIndex]=arr[swapIndex],arr[pivotIndex]
	return swapIndex
}

func QuickSortHelper(arr []int ,left,right int){
	if left < right{
		pivotIndex:=Pivot(arr,left,right)
		QuickSortHelper(arr,left,pivotIndex-1)
		QuickSortHelper(arr,pivotIndex+1,right)
	}
}

func QuickSort(arr []int){
	QuickSortHelper(arr,0,len(arr)-1)
}