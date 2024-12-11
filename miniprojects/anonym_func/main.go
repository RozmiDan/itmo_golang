// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var val uint
// 	fmt.Scan(&val)

// 	fn := func(val uint) uint {
// 		res := make([]rune, 0, 4)
// 		str := strconv.Itoa(int(val))
// 		for _, val := range(str) {
// 			if cur, _ := strconv.Atoi(string(val)); cur % 2 == 0 && cur != 0{
// 				if(cur == 0 && len(res) == 0) {
// 					continue
// 				}
// 				res = append(res, val)
// 			}
// 		}
// 		if(len(res) == 0 || res[0] == 0) {
// 			return 100
// 		} else {
// 			val, _ := strconv.Atoi(string(res))
// 			return uint(val)
// 		}
// 	}
// 	fmt.Println(fn(val))
// }

package main

import "fmt"

func readTask() (interface{}, interface{},interface{}) {
	var b float64
	var a int
	var c string
	fmt.Scan(&a, &b, &c)
	return a, b, c
}

func check(value1, value2, operation interface{}){
	flag := true
	var arg1, arg2 float64
	switch value1.(type){
	case float64: 
		arg1 = value1.(float64)
	default: 
		flag = false
		fmt.Printf("value=%d: %T", value1, value1)
	}
	if(flag){
		switch value2.(type){
		case float64: 
			arg2 = value2.(float64)
		default: 
			flag = false
			fmt.Printf("value=%d: %T", value2, value2)
		}
	}
	if(flag){
		switch operation.(type){
		case (string):
			if(operation == "*"){
				fmt.Printf("%.4f", arg1 * arg2)
			} else if(operation == "/") {
				fmt.Printf("%.4f", arg1 / arg2)
			} else if(operation == "+") {
				fmt.Printf("%.4f", arg1 + arg2)
			} else if(operation == "-") {
				fmt.Printf("%.4f", arg1 - arg2)
			} else {
				fmt.Print("неизвестная операция")
			}
		}
	}
	return
}

func main(){
	value1, value2, operation := readTask()
	check(value1, value2, operation)
	//fmt.Print(value1, value2, operation)
}