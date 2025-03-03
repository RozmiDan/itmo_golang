package main

import (
	"fmt"

	"github.com/RozmiDan/storage/pkg/storage"
)

func main() {
	st := storage.NewStorage()
	file , err := st.Upload("test.txt", []byte("helo"))
	if err != nil {
		panic("error")
	}
	fmt.Println(file)
}