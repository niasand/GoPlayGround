package GoExamplesReadOnly

import "fmt"

func mainfalse()  {

	var a int = 4
	//var b int32
	//var c float32
	var pointer *int


	/*  & 和 * 运算符实例 */

	/*
		&  返回变量存储地址	&a; 将给出变量的实际地址。
		*  指针变量。         *a; 是一个指针变量

	 */

	pointer = &a  /* 'pointer' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a);
	fmt.Println("pointer的实际地址",pointer)
	fmt.Printf("pointer的值为 %d\n",*pointer)

}
