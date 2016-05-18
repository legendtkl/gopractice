package main

import (
	"fmt"
	"test/ffjson"
)

func main() {
	user := &ffjson.User{
		Id:        123,
		UserName:  "taokelu",
		CellPhone: "15652190000",
	}

	res, _ := user.MarshalJSON()

	fmt.Println(string(res))

}
