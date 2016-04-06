package main

import (
	"fmt"
	//	"runtime"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	go func() {
		fmt.Println("channel data: ", <-ch)
	}()
	//	runtime.Gosched()
}

/*
func main() {
	ch := make(chan int, 5)
	sign := make(chan byte, 5)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
		close(ch)
		fmt.Println("the channel is closed")
		sign <- 0
	}()
	go func() {
		for {
			e, ok := <-ch
			fmt.Printf("%d (%v)\n", e, ok)
			if !ok {
				break
			}
			time.Sleep(2 * time.Second)
		}
		fmt.Println("Done.")
		sign <- 1
	}()
	<-sign
	<-sign
}*/
