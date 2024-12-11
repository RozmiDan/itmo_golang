package main

import (
	"bufio"
	"encoding/json"
	"os"
)

type PersonInfo struct {
	ID int 
	Number string
	Year int
	Students []Stud
}

type Stud struct{
	Rating []int
}

type Result struct{
	Average float32
}

func main() {
	// scaner := bufio.NewScanner(os.Stdin)
	// scaner.Scan()
	// data := scaner.Text()
	// os.Stdout.WriteString(string(data))

	// fileName := "info.json"
	// file, err := os.Open(fileName)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// data, err := os.ReadFile(fileName)
	scanner := bufio.NewScanner(os.Stdin)
  var data string
	for scanner.Scan(){
		value := scanner.Text()
		if len(value) == 0 { 
			break
		}
		data += value
	}
	if err := scanner.Err(); err != nil {
		panic(err)	
  }
	var pers PersonInfo
	if err := json.Unmarshal([]byte(data), &pers); err != nil {
		panic(err)
	} 
	var ratingCount int
    for _, student := range pers.Students {
        ratingCount += len(student.Rating)
    }
	res := Result{float32(ratingCount)/float32(len(pers.Students))}
	str_answer, err := json.MarshalIndent(res, "" ,"    ")
	if err != nil{
		panic(err)
	}
	os.Stdout.WriteString(string(str_answer))
}