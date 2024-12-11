package main

import (
	"fmt"
	"math"
)

var b int = 23

func main() {
	var k,p,v float64
	fmt.Scan(&k, &p, &v)
	m := M(p, v)
	w := W(k, m)
	t := T(w)
	fmt.Println(t)
}

func M(p, v float64) float64 {
	m := p * v
	return m
}

func W(k, m float64) float64 {
	w := math.Sqrt(k / m)
	return w
}

func T(w float64) float64 {
	t := 6 / w
	return t
}