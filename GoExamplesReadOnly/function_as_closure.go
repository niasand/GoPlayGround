/*
* @Author: jerry
* @Date:   2017-09-21 17:06:31
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-21 17:16:56
 */

package GoExamplesReadOnly

import "fmt"

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func mainfalse() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

}
