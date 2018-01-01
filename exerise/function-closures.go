package main

import (
	"fmt"
)

func adder() func(int) float64 {
	sum := 0
	return func(x int) float64 {
		sum += x
		return float64(sum) * 1.22
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
