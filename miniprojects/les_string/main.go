package main

import (
	"fmt"
)

// Ex 1

// func main() {
// 	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
// 	rune_text := []rune(text)
// 	//fmt.Println(rune_text)
// 	sizeOfStr := utf8.RuneCountInString(text)

// 	if unicode.IsUpper(rune_text[0]) && rune_text[sizeOfStr-3] == '.'  {
// 		fmt.Println("Right")
// 		//fmt.Printf("%c", text[sizeOfStr-3])
// 	} else {
// 		fmt.Println("Wrong")
// 	}

// 	// en := "Быть или не быть."
// 	// ru := " русский "
// 	// fmt.Println(len(en), len(ru))
// 	// s := utf8.RuneCountInString(en)
// 	// fmt.Printf("%c", en[s-1])
// 	// if en[s-1] == 46 {
// 	// 	print("gfgdfgdfgfd")
// 	// }
// }

// Ex 2
// func main() {
// 	var text string
// 	var flag bool = true
// 	fmt.Scan(&text)
// 	rune_text := []rune(text)
// 	lstIndStr := utf8.RuneCountInString(text) - 1
// 	for i:=0; i<lstIndStr; i++ {
// 		if rune_text[i] != rune_text[lstIndStr] {
// 			flag = false
// 			break
// 		}
// 		lstIndStr--
// 	}

// 	if flag {
// 		fmt.Println("Палиндром")
// 	} else {
// 		fmt.Println("Нет")
// 	}
// }

// //Ex3
// func main() {
// 	var text string
// 	fmt.Scan(&text)
// 	rune_text := []rune(text)
// 	var r rune = '*'
// 	szText := utf8.RuneCountInString(text)
// 	res := make([]rune, 1, szText * 2)
// 	res[0] = rune_text[0]
// 	//fmt.Printf("%c", res[0])
// 	for i:=1; i<szText-1; i++ {
// 		res = append(res, r, rune_text[i])
// 	}
// 	res = append(res, r, rune_text[szText-1])
// 	text = string(res)
// 	fmt.Println(text)
// }

//Ex4
func main() {
	var text string
	fmt.Scan(&text)
	var res rune
	rune_text := []rune(text)
	for i:=0; i < len(rune_text); i++ {
		// fmt.Println(rune_text[i] - '0')
		// fmt.Println(res)
		if res < (rune_text[i]) {
			res = rune_text[i]
		}
	}
	fmt.Printf("%c", res)
}