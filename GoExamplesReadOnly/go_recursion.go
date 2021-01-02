/*
* @Author: jerry
* @Date:   2017-09-22 11:49:16
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-22 12:00:16
 */
/*递归，就是在运行的过程中调用自己

Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。

func recursion() {
   recursion()
}

func mainfalse() {
   recursion()
}
*/
package GoExamplesReadOnly

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {

		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)

}

func mainfalse() {
	var i uint64 = 3
	fmt.Printf("%d 的阶乘是%d\n", i, Factorial(i))
	var n int
	for n = 1; n < 10; n++ {
		fmt.Printf("%d\n", Fibonacci(n))
	}

}
