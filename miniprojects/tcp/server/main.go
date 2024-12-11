package main

import (
	"fmt"
	"net"
)

func requestHandler(r *net.Listener, w){

}

func main(){
	listener, err := net.Listen("tcp", "8080")
	if(err != nil){
		fmt.Println("bad gateway")
		panic("xx")
	}
	newCon, err := listener.Accept()
	if(err != nil){
		panic("cant connect")
	}
	

	defer listener.Close()

}