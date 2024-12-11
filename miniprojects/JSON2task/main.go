package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type(
	// Arr struct{
	// 	Vec []Struc
	// }

	Struc struct{
		Global_id int64 `json:"global_id"`
	}
)

func main() {
	fileName := "data-20190514T0100.json"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var resJson []Struc
	if err := json.Unmarshal(data, &resJson); err != nil {
		panic(err)
	} 
	var sumOfId int64
	for _, val := range resJson {
		sumOfId += val.Global_id
	}

	fmt.Print(sumOfId)

}