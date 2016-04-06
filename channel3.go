package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)

	<-done
	fmt.Println("Done!")
}

func main() {
	done := make(chan bool)
	go worker(done)
	fmt.Print("Before")
	done <- true
	fmt.Println("After")
}
