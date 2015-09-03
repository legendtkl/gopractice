package main

import (
	"os"
	"syscall"
)

func main() {
	oldMask := syscall.Umask(0)
	os.Mkdir("abc", os.ModePerm)
	syscall.Umask(oldMask)
}
