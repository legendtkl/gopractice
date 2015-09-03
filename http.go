package main

import "net/http"
import "fmt"

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		return
	} else {
		fmt.Println(resp)
	}
}
