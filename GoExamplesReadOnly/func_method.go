/*
* @Author: jerry
* @Date:   2017-09-21 17:39:55
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-21 17:45:48
 */
package GoExamplesReadOnly

import (
	"fmt"
)

/*
Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体
类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：

func (variable_name variable_data_type) function_name() [return_type]{
   /* 函数体*/

// 定义一个结构体类型
type Circle struct {
	radius float64
}

func mainfalse() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
}

func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}
