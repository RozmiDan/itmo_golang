// package main

// import "fmt"

// func main() {
// 	// var a float32 = 10
// 	// var b = &a
// 	// var c = b
// 	// var q,w,e string
// 	// fmt.Println(q,w,e)
// 	// const j float32 = 12.4
// 	// //x int = 54
// 	// //fmt.Println(x) - error
// 	// fmt.Println(j + 23)
// 	// *c += 2
// 	// fmt.Println(a + *b)
// 	// fmt.Println(*c)

// 	// const(
// 	// 	a = 10 + iota
// 	// 	_
// 	// 	c
// 	// 	d
// 	// 	e
// 	// )
// 	// fmt.Print(a, c, d, e)

// 	// var val int32
// 	// fmt.Scan(&val)
// 	// if val > 0 {
// 	// 		fmt.Println("Число положительное")
// 	// } else if val < 0 {
// 	// 		fmt.Println("Число отрицательное")
// 	// } else {
// 	// 		fmt.Println("Ноль")
// 	// }

// }

// package main

// import "fmt"
// func main(){

//   var a, b int
//   fmt.Scan(&a) // считаем переменную 'a' с консоли
//   fmt.Scan(&b) // считаем переменную 'b' с консоли

//   a = a * a
//   b = b * b
//   var c = a + b
//   fmt.Println(c)
// }

// package main

// import "fmt"

// func main() {
//     var fst, scnd uint8
//     fmt.Scan(&fst, &scnd)
// 		fmt.Println(fst + scnd)
// }

// package main

// import "fmt"

// func main() {
//     var fst, scnd int
//     var result uint16
//     fmt.Scan(&fst, &scnd)
//     for ; fst <= scnd; fst++ {
//         result += uint16(fst)
//     }
// 		fmt.Println(result)
// }

package main

import "fmt"

func main() {
    var str string
    fmt.Scan(&str)
    // runeStrF := []rune(str)
    for _, val := range(str) {
		if val >= '0' && val <= '9' {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
