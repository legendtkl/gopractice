package main

import (
	"fmt"
	"runtime"
)

func main() {
	/*
		names := []string{"Eric", "Harry", "Robert", "Jim", "Mark"}
		for _, name := range names {
			for i := 1; i < 10; i++ {
				go func() {
					fmt.Printf("Hello, %s.\n", name)
				}()
			}
			runtime.Gosched()
		}*/
	println(runtime.GOMAXPROCS())
	fmt.Println(runtime.NumGoroutine())
}
