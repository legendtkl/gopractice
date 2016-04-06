package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	file := "./hello/world/legendtkl"
	fmt.Println(filepath.Base(file))
}
