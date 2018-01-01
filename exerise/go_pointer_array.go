/*
* @Author: jerry
* @Date:   2017-09-21 19:09:33
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-21 19:13:28
 */
package main

import "fmt"

const (
	MAX int = 3
)

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
		fmt.Println(ptr, ptr[i])
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d]= %d\n", i, *ptr[i])

	}
}
