package main

import (
	"fmt"
)

func incr(p *int) int {
	*p++   // 非常重要：只是增加p指向的变量的值，并不改变p指针！！
	return *p

}

func main() {
	v := 1
	incr(&v)
	fmt.Println(incr(&v))




}
