package main

import "fmt"

func main() {
	// a := [...] int {2,7,8,23}
	// for ind, val := range a {
	// 	fmt.Println(ind, val)
	// }
	// fmt.Println("Slice:")
	// var sliceA [] int = a[:]
	// sliceB := a[2:4]
	// sliceC := make([]int, 0, 4)
	// sliceC = append(sliceC, sliceB...)
	// fmt.Println(sliceA)
	// fmt.Println(sliceB)
	// fmt.Println(sliceC)

// 	a := [3] int {23,54,12}
// //	fmt.Println("Adr a: ", &a[0])
// 	b := a[:]
// 	// fmt.Println("fst operation: ", cap(b), len(b))
// 	// fmt.Println("adr b do: ", &b[0])
// 	b = append(b, a[:]...)	
// 	// fmt.Println("scnd operation: ",cap(b), len(b))
// 	fmt.Println("Adr a: ", &a[0])
// 	fmt.Println("adr b posle: ", &b[0])
// 	sl_a := make([]int, 2, 7)
// 	sl_a  = b[1:2]
// 	var max int = -100000000
// 	fmt.Println("b: ", b, len(b), cap(b), "adr b ", &b[1])
// 	fmt.Println("sl_a: ", sl_a, len(sl_a), cap(sl_a), "adr sl_a ", &sl_a[0])
// 	fmt.Printf("%e", max)


	// var N int
	// fmt.Scan(&N)
	// sl_a := make([]int, N, N)
	// for i:=0; i<N; i++ {
	// 		fmt.Scan(&sl_a[i])
	// 		if i % 2 == 0 {
	// 				fmt.Printf("%d ", sl_a[i])
	// 		}
	// }

	var N, cntr int
	fmt.Scan(&N)
	var arr []int = make([]int, N, N)
	for i:=0; i<N; i++ {
			fmt.Scan(&arr[i])
			if arr[i] >= 0 {
					cntr++
			}
	}
	fmt.Println(cntr)

}