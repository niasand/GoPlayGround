package main

import (
	"fmt"
)

func annonymous() func(x, y int) int {
	var m int = 2
	var n int = 1
	return func(x, y int) int {
		return x + y + m + n
	}
}

func main() {
	f := annonymous()
	fmt.Println(f(3, 4))
}
