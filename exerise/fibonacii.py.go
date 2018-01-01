package main

import "fmt"

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Println(i, result)

	}
}

func fibonacci(n int) int {
	var res int
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}

	return res
}

func print_ten_to_zero(n int ) int {
	var j int
	if n <=1{
		j = 1
	}
	else{
		j -=1
		fmt.println

	}
}

