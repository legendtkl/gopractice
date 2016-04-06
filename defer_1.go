package main

import "fmt"

func PrintNum(x int) {
	fmt.Println(x)
}

func main() {
	x := 5
	defer PrintNum(x)
	x += 5
}
