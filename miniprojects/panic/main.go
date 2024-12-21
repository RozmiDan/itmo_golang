package main

import (
	"fmt"
)

func checker(checkFunc func()){
	defer func(){
		if val := recover(); val != nil{
			fmt.Printf("runtime panic catched: %s\n", val)
		}
	}()

	fmt.Println("func start")
	checkFunc()
	fmt.Println("func end")
}

func main(){
	checker(func(){
		panic("random err")
	})
	fmt.Println("programm end")
}