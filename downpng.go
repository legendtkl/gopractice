package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	s := "http://m.360buyimg.com/n4/jfs/t730/218/249574523/86833/7128dabd/549930a8N95eee720.jpg!q70.jpg"
	ret := strings.Split(s, "/")
	n := len(ret)
	resp, _ := http.Get(s)
	filename := ret[n-1]
	fout, _ := os.Create(filename)
	io.Copy(fout, resp.Body)
	/*
		fmt.Println(ret[-1])
			for _, v := range strings.Split(s, "/") {
				fmt.Println(v)
			}*/
}
