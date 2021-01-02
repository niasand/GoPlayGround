package GoExamplesReadOnly

import "fmt"

func fibonacci(n int) (res int) {
	if n < 2 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}

func print10(n int) (res int) {
	if n < 2 {
		fmt.Println(n)
	} else {
		fmt.Println(print10(n - 1))
	}
	return n + 1

}

func mainfalse() {
	print10(10)
	result := 0
	for i := 0; i < 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is %d\n", i, result)
	}
}
