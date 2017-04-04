package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	const world = "世界"
	fmt.Println("Hello ", world)
	fmt.Println("Happy", Pi, "Day")

	const (
		Truth = true
		// No    = false
	)
	fmt.Println("Go Rules?", Truth)
}

/*

常量
常量的定义与变量类似，只不过使用 const 关键字。

常量可以是字符、字符串、布尔或数字类型的值。

常量不能使用 := 语法定义。
*/
