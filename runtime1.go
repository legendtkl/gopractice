package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(2))
	fmt.Println(runtime.NumGoroutine())
}
