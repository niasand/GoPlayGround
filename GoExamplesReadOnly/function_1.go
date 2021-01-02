/*
* @Author: jerry
* @Date:   2017-09-21 16:17:24
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-21 16:43:00
 */
package GoExamplesReadOnly

import "fmt"

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

func swap_string(x, y string) (string, string) {
	return y, x
}

func mainfalse() {
	var num1, num2 = 1, 3
	var r int
	r = max(num1, num2)
	fmt.Printf("max is %d\n", r)

	a, b := swap_string("zhiwei", "yang")
	fmt.Println(a, b)
}
