package main

import (
	//	"io"
	//	"io/ioutil"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
	logger.Println("This is a DEBUG LOG")
}
