package random

import "math/rand"

func NewAliasForURL(lengthOfStr int) string {
	var arrOfNums = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

	result := make([]rune, lengthOfStr)
	for i := 0; i < lengthOfStr; i++ {
		result[i] = arrOfNums[rand.Intn(len(arrOfNums))]
	}

	return string(result)
}
