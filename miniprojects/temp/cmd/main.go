package main

import (
	"fmt"
	"unsafe"
)

type Comparablya interface{
	rune
}

func mapChecker[K Comparablya, V comparable](m map[V]K) {
	fmt.Println(len(m))
}

func main() {
	// mp := make(map[int]string)
	// mp[2] = "sdfds"
	// mp[12] = "fd"
	// var val string = mp[10]
	// fmt.Println(val)
	// if v,ok := mp[3]; ok{
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println(len(mp))
	// }
	// nmp := make(map[float32]rune)
	// nmp[2.43] = 'f'
	// mapChecker(nmp)
	var arr [4]int = [4]int{1,3,4,5}
	sl := arr[0:4]
	fmt.Println(&arr[0])
	fmt.Println(unsafe.Pointer(&sl[0]))
	fmt.Println(sl)
}