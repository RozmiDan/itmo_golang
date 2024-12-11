package main

import "fmt"

func main() {
	testStruct := new (myStruct)
	testStruct.Ammo = 1
	testStruct.On = true
	testStruct.Power = 1
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.Shoot())
	fmt.Println(testStruct.RideBike())
}