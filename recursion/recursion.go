package recursion

import "fmt"

func Factorial(n int) int{
	if n<=1{
		return 1
	}
	return n*Factorial(n-1)
}

func Sample(n int){
	if n<=1{
		return
	}
	Sample(n-1)
	fmt.Println(n)
	Sample(n-1)
}

func FibonacciSeries(n int) int{
	if n<=1{
		return n
	}
	return FibonacciSeries(n-1)+FibonacciSeries(n-2)
}