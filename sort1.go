package main

import (
	"fmt"
	//	"sort"
)

func (A []int) Len() int { return len(A) }

func main() {
	A := []int{3, 1, 2, 5, 4}
	//	sort.Sort(A)
	fmt.Println(Len(A))
}
