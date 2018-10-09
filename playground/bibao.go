package main

import "fmt"

func Adder2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func main() {
	p2 := Adder2()
	fmt.Printf("CAll add2 for 3 givens: %v\n", p2(3))

}
