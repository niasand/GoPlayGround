package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprintln(math.Sqrt(x))
}

func get_negtive_number(x int) int {
	if x > 0 {
		return (-x)
	}
	if x == 0 {
		return 888
	}
	return x
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(get_negtive_number(2), get_negtive_number(-4), get_negtive_number(0))
}

// if就像 for 循环一样，Go 的 if 语句也不要求用 ( ) 将条件括起来，同时， { } 还是必须有的。
