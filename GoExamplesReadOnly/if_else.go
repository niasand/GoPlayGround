package GoExamplesReadOnly

import "fmt"

func mainfalse() {
	var a int = 100
	var b int = 10

	if a < 20 {
		fmt.Printf("a 小于 20 ")
	} else {
		fmt.Printf("a 不小于 20 ")
	}
	fmt.Printf("a的值为: %d\n", a)

	if a == 100 {
		if b == 10 {
			fmt.Printf("a的值为%d，b的值为%d\n", a, b)
		}
	}
}
