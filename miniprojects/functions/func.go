package main

func sumInt(arr_args ...int) (int, int) {
	var sum int
	quan_args := len(arr_args)
	for _, elem := range arr_args {
		sum += elem
	}
	return quan_args, sum
}