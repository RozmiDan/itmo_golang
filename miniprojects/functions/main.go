package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	arr := make([]int, N)
	for i:=0; i<N; i++ {
		fmt.Scan(&arr[i])
	}
	q, s := sumInt(arr...)
	fmt.Println(q, s)
}


