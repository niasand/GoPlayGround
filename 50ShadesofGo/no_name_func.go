package main

import "fmt"

func main() {
	//匿名函数1
	//f1 为函数地址
	f1 := func(x, y int) (z int) {
		z = x + y
		return
	}
	fmt.Println(f1)
	fmt.Println(f1(5, 6))

	//匿名函数2
	// 直接创建函数并且运行
	f2 := func(x, y int) (z int) {
		z = x + y
		return
	}(7, 8)
	fmt.Println(f2)

	//匿名函数2 无参数的形式， 也是创建函数后直接运行
	func() {
		fmt.Println(9 + 10)
	}()
}
