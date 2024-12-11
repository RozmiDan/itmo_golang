package main

import (
	utils "concarencyFst/myModuls"
	"fmt"
)

func main() {
    result := utils.Sum(3, 4)
    fmt.Println("Сумма:", result)

    product := utils.Multiply(3, 4)
    fmt.Println("Произведение:", product)
}