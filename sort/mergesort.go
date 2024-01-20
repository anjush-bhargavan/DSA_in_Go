package sort


func merge(arr1,arr2 []int) []int {
	combined:=make([]int,len(arr1)+len(arr2))
	index,i,j:=0,0,0
	for i<len(arr1) && j <len(arr2){
		if arr1[i]<arr2[j]{
			combined[index]=arr1[i]
			index++
			i++
		}else{
			combined[index]=arr2[j]
			index++
			j++
		}
	}
	for i<len(arr1){
		combined[index]=arr1[i]
		index++
		i++
	}
	for j<len(arr2){
		combined[index]=arr2[j]
		index++
		j++
	}
	return combined
}


func MergeSort(arr []int) []int {
	if len(arr)==1{
		return arr
	}
	mid:=len(arr)/2

	left:=MergeSort(arr[:mid])
	right:=MergeSort(arr[mid:])

	return merge(left,right)
}