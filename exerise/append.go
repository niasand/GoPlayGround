package main

import (
	"fmt"
)

func main() {
	// var a []int
	a := make([]int, 5, 8)
	printSlice("a", a)

	// append works on nil needed
	a = append(a, 1)
	printSlice("a", a)

	//we can ad more than one element at a time
	a = append(a, 2, 3, 4)
	printSlice("a", a)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
