package main

import (
	//"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
//	fileName := "test.txt"
	// dataForWrite := "this string is in file 1"
	// os.WriteFile(fileName, []byte(dataForWrite), 0600)
	// os.WriteFile(fileName, []byte(dataForWrite + "23"), 0600)
	// if res, err := os.ReadFile("test1.txt"); err != nil {
	// 	fmt.Print("error happend")
	// } else {
	// 	fmt.Print(string(res))
	// }
	// file, err := os.OpenFile(fileName, os.O_APPEND, 0600)
	// if err != nil {
	// 	panic(err)
	// } 
	// defer file.Close()
	// file.WriteString("Some new string")
	// fmt.Print()
	res_sum := 0
	in := os.Stdin
	out := os.Stdout
	scaner := bufio.NewScanner(in)
	for scaner.Scan(){
		value := scaner.Text()
		if len(value) == 0 { 
			break
		}
		var val_conv int
		var err error
		if val_conv, err = strconv.Atoi(value); err != nil{
			panic(err)
		}
		res_sum += val_conv
	}
	out.WriteString(strconv.Itoa(res_sum))
}