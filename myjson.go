package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "hello", 12345678}
	b, _ := json.Marshal(m)
	fmt.Println(string(b))

	b = []byte(`{"Name":"Bob", "Food":"Pickle"}`)
	msg := &Message{}
	if err := json.Unmarshal(b, msg); err != nil {
		fmt.Println("Error")
	}
	fmt.Println(string(msg.Name))
	fmt.Println(string(msg.Body))
}
