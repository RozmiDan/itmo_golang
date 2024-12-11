package main

import "fmt"

type Inter struct{
	l int
	s string
}

func (inter* Inter) Unter(r int)  {
	inter.l = r
}

func main() {
	c := new(Inter)
	c.Unter(23)
	fmt.Println(c.l)
}