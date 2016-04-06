package main

import "fmt"
import "strconv"

func main() {
	s := "12345"
	a := 3
	t := strconv.Itoa(a)
	s += t
	fmt.Println(s[2:4])
	fmt.Println(t)
}
