package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5) //len(a) is?
	fmt.Println(a, len(a))

	b := make([]int, 3, 6)

	fmt.Println(b, len(b))
}
