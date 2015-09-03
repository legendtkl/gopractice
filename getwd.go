package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	var currentDir string
	currentDir, _ = os.Getwd()
	fmt.Println("current dir = %s", currentDir)
	fmt.Println(path.Base(currentDir))
	fmt.Println(path.Dir(currentDir))
	fmt.Println(path.IsAbs(currentDir))
	_, s := path.Split(currentDir)
	fmt.Println(s)
}
