package sort


func InsertionSort(arr []int){
	for i:=1;i<len(arr);i++{
		temp:=arr[i]
		j:=i-1
		for j >= 0 && temp < arr[j] {
			arr[j+1],arr[j]=arr[j],temp
			j--
		}
	}
}