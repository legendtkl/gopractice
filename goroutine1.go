package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 0; i < 100000; i++ {
		go func(num int) {
			fmt.Printf("Hello, %d.\n", num)
		}(i)
	}
	runtime.Gosched()
}
