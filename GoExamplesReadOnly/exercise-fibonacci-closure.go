package GoExamplesReadOnly

import (
	"fmt"
)

//fibonacci 函数会返回一个返回 int 的函数。

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func mainfalse() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
