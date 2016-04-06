package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Num  int
}

func (p *Person) Init() *Person {
	p.Name = "kelu"
	p.Age = 19
	return p
}
func main() {
	//p := new(Person).Init()
	p := new(Person)
	fmt.Println(*p)
}
