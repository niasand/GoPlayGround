/*
* @Author: jerry
* @Date:   2017-09-23 17:15:01
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-23 18:23:25
 */
package main

import (
	"fmt"
)

func main() {
	i := 5
	fmt.Printf("A integer: %d,it is location memory: %p\n", i, &i)
	/*
			   程序 string_pointer.go 为我们展示了指针对string的例子。

			   它展示了分配一个新的值给 *p 并且更改这个变量自己的值（这里是一个字符串）。

		        通过对 *p 赋另一个值来更改“对象”，这样 s 也会随之更改。

	*/
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("here is the pointer p: %p\n", p)
	fmt.Printf("here is the string  *p: %s\n", *p)
	fmt.Printf("here is the string  &p: %p\n", &p)
	fmt.Printf("here is the string  s: %s\n", s)

}
