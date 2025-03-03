package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type MyClass struct {
	fstVal  int			`json:"renamedFstValue"`
	scndVal string	`json:"-"`
	thrdVal string	`json:"mapStruct, omitempty"`
}

func (m *MyClass) MarshalJSON() ([]byte, error) {
	const prefix, suffix = `{"privateField": "`, `"}`

	buf := new(bytes.Buffer)
	buf.Grow(len(prefix) + len(m.scndVal) + len(m.thrdVal) + len(suffix))

	buf.WriteString(prefix)
	buf.WriteString(strconv.Itoa(m.fstVal))
	buf.WriteString(m.thrdVal)
	buf.WriteString(suffix)
	return buf.Bytes(), nil
}

var marshI json.Marshaler = &MyClass{}

func main() {
	data := &MyClass{
		fstVal: 23,
		scndVal: "hello",
		thrdVal: "world",
	}

	bufer, err := json.Marshal(data)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(bufer))
}