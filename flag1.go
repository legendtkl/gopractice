package main

import (
	"flag"
	"fmt"
)

func main() {
	var version *bool = flag.Bool("v", false, "version")
	flag.Parse()

	fmt.Println(*version)
	fmt.Println(flag.NArg())
	fmt.Println(flag.Args())
}
