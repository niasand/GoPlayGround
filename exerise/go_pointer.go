/*
Go 语言中指针是很容易学习的，Go 语言中使用指针可以更简单的执行一些任务。
接下来让我们来一步步学习 Go 语言指针。
我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
以下实例演示了变量在内存中地址：
*/

package main

import "fmt"

func main() {
	var a int = 10 /* 声明实际变量 */
	fmt.Printf("a变量的地址: %x\n", &a)

	var ip *int /* 声明指针变量 */
	var ptr *int
	ptr = &a
	if ptr == nil {
		fmt.Printf("Nil ptr value is %x\n", ptr)
	} else {
		fmt.Printf("指针存储的变量地址: %x\n", ptr)
	}

	ip = &a /* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip变量值:%d\n", *ip)

}

/*
如何使用指针
指针使用流程：
定义指针变量。
为指针变量赋值。
访问指针变量中指向地址的值。
在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
*/
