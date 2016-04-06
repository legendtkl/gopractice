package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return "hello world"
}

func main() {
	people := Person{"Bob", 31}
	fmt.Println(people)
}
