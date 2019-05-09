package main

import "fmt"

//Shape 定义一个Shape的接口，可以对不同形状的shape求面积
type Shape interface {
	area() int
}

//Square 一个矩形的结构体，求面积就是长乘以宽
type Square struct {
	width  int
	height int
}

// 正方形面积的计算方法
func (s Square) area() int {
	return s.height * s.width
}

//  统一在一起，计算某种形状的面积
func thearea(s Shape) int {
	return s.area()
}

func main() {
	s := Square{3, 4}
	v := thearea(&s)
	fmt.Println(v)
}
