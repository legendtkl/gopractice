package main

import (
	"fmt"
)

type MyInt int

func (f MyInt) square() int {
	return int(f) * int(f)
}

func main() {
	f := MyInt(4)
	fmt.Println(f.square())
}
