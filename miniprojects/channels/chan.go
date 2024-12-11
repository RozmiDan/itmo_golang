package main

import (
	"fmt"
)

func main(){
	
	fst_chan := make(chan int, 2)
	
	fun_fst := func(val interface{}){
		if int_val, ok := val.(int); ok {
			fst_chan <- int_val
			fmt.Println("fst func write:", int_val)
		} else {
			fst_chan <- 52
			fmt.Println("fst func write:", 52)
		}
		fmt.Println("fst func end")
	}
	
	fun := func(){
		for i := range fst_chan {
			fmt.Println("scnd func read:", i)
		}
		fmt.Println("func end")
	}

	go fun()
	go fun_fst(23)
	go fun_fst("lox")
	go fun_fst(23.3)
	go fun_fst(423)
	go fun_fst(4)
	go fun_fst(12)

	fmt.Scanln()

}

// package main

// import (
// 	"fmt"
// )

// func main(){
	
// 	data_chan := make(chan int)
// 	cancel_chan := make(chan error)

	
// 	go func(dt_ch chan<- int, cncl_ch <-chan error){
// 		val := 0
// 		for {
// 			select{
// 			case <-cncl_ch:
// 				return
// 			case dt_ch<-val:
// 				val++
// 			default:
// 				fmt.Println("default")
// 			}
// 		}
// 	}(data_chan, cancel_chan)
	
// 	for{
// 		i :=  <-data_chan 
// 		fmt.Println("read:", i)
// 		if(i > 6){
// 			var b error 
// 			cancel_chan<-b
// 			fmt.Println("Close datach")
// 			break
// 		}
// 	}
// 	fmt.Println("func end")
// 	fmt.Scanln()

// }