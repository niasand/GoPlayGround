package GoExamplesReadOnly

import "fmt"

func mainfalse() {
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
