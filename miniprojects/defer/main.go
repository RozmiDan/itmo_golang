package main

import (
	"fmt"
	"sync/atomic"
)

var a string
var flag atomic.Bool

func setup(){
	a = "hello"
	flag.Store(true)
}

func main(){
	go setup()
	for !flag.Load(){}
	fmt.Print(a)
}

// func main(){
// 	i := 0
// 	mut := new(sync.Mutex)
// 	go func(){
// 		mut.Lock()
// 		i++
// 		mut.Unlock()
// 	}()
// 	mut.Lock()
// 	fmt.Println(i)
// 	mut.Unlock()
// }

// Fst example of using defer

// type MyInterface interface{
// 	Foo()
// }

// func main(){
// 	var l MyInterface
// 	a := 0

// 	defer func(){
// 		l.Foo()
// 	}()

// 	a++
// 	fmt.Println(a)
// }