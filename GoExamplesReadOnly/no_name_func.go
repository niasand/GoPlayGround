package GoExamplesReadOnly

import "fmt"

func anoymous() {
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
	}(7, 8) //花括号后加()表示函数调用，，并将返回的函数指针赋给变量f2
	fmt.Println(f2)

	//匿名函数2 无参数的形式， 也是创建函数后直接运行
	func() {
		fmt.Println(9 + 10)
	}() //花括号后加()表示函数调用

	fn := func() {
		fmt.Println("hello")
	}
	fmt.Printf("%T\n", fn) //打印fn的类型
	//我们现在有一个变量fn，它是一个function；它的类型是func()。
	//它能像其他的任何函数一样被调用，通过表达fn(),或者赋值给你感兴趣的其他func()。

	fn1 := func(a, b int, z float64) bool {
		return a > b
	}
	fmt.Printf("%T\n", fn1) //打印fn1的类型,   即func(int, int, float64) bool

	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2) //直接调用

	//另一种写法
	fn2 := func(a, b int) {
		fmt.Println(a + b)
	}
	fn2(2, 3) //把匿名函数这 复制给fn2，那么fn2这个时候他的类型就是函数，因为现在fn2是个函数了，
	//所以你可以把fn2当作函数来使用，该传参就传参数，该不传参数就不传参数
}
