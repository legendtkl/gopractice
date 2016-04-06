package main

import (
	"fmt"
	"log"
	//	"os"
)

func main() {
	/*
		f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("file open error : %v", err)
		}
		defer f.Close()
		//log.SetOutput(f)
		log.SetPrefix("legendtkl")
		log.SetFlags(0)

		log.Println("This is a test log entry 1")
		log.Println("This is a test log entry 2")
	*/
	fmt.Println(log.Ldate)
	fmt.Println(log.Ltime)
	fmt.Println(log.Lmicroseconds)
}
