package sort

func SelectionSort(arr []int){
	for i:=0;i<len(arr)-1;i++{
		minIndex:=i
		for j:=i+1;j<len(arr);j++{
			if arr[minIndex] > arr[j]{
				minIndex=j
			}
		}
		if i!=minIndex{
			arr[i],arr[minIndex]=arr[minIndex],arr[i]
		}
		
	}
}