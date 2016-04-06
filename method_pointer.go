package main

import "fmt"

func func_value(num ...int) int {
	num[0] = 1
	return len(num)
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	n := func_value(a...)
	fmt.Println(n, a)
}
