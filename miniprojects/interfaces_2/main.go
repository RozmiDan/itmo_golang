package main

import (
	"fmt"
)

type Batery struct{
	Power string
}

func (bat Batery) String() string{
	return fmt.Sprintf("[%s]", bat.Power)
}

func init_s (str *string) *Batery{
	res := make([]rune, 0, 10)
	counter := 0
	for _, val := range []rune(*str) {
		if(val == '1'){
			counter++
		}
	}
	for i:=0; i < 10; i++ {
		if(i + counter >= 10) {
			res = append(res, 'X') 
		} else {
			res = append(res, ' ') 
		}
	}
	return &Batery{string(res)}
	// fmt.Print(string(res))
}

func main(){
	var val string 
	fmt.Scan(&val)
	var batteryForTest  = init_s(&val)
	fmt.Print(batteryForTest )
}