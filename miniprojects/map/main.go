package main

import "fmt"

func main() {
	const n int = 10
	var arr [10]int
	map_a := make(map[int]int)
	for i:=0; i<n; i++ {
		fmt.Scan(&arr[i])
	}

	for i:=0; i<n; i++ {
		if val_v, ok := map_a[arr[i]]; ok {
			arr[i] = val_v
			continue
		} else {
			res := work(arr[i])
			map_a[arr[i]] = res
			arr[i] = res
		}
	} 
	for _, elem := range(arr) {
		fmt.Printf("%d ", elem)	
	}
}

func work(x int) int {
	if x > 3 {
		return x + 1
	} else {
		return x - 1
	}
}