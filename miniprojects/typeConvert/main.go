package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var fstVal, scndVal string
	fmt.Scan(&fstVal, &scndVal)
	res, ok := adding(fstVal, scndVal)
	if ok != nil {
		panic("Can't sum!")
	}
	fmt.Println(res)
}

func adding(strFst, strScnd string)(res int64, err error) {
	strFltrFst, strFltScnd := filter(strFst, strScnd)
	valFst, ok := strconv.Atoi(strFltrFst)
	if ok != nil {
		err = errors.New("Error in converion first operand") 
	}
	valScnd, ok := strconv.Atoi(strFltScnd)
	if ok != nil {
		err = errors.New("Error in converion second operand")
	}
	res = int64(valFst + valScnd)
	return res, err
}

func filter(fstStr, scndStr string)(string ,string) {
	var resFst, resScnd = make([]rune, 0), make([]rune, 0) 
	for _, val := range(fstStr) {
		if val >= '0' && val <= '9' {
			resFst = append(resFst, val)
		}
	}
	for _, val := range(scndStr) {
		if val >= '0' && val <= '9' {
			resScnd = append(resScnd, val)
		}
	}
	return string(resFst), string(resScnd)   
}