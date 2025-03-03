package main

import (
	"reflect"
)

func main(){
	var a = 4
	var b = "4"
	reflect.DeepEqual(a,b)

}