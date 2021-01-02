/*
* @Author: jerry
* @Date:   2017-09-21 17:46:49
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-21 17:54:28
 */
package GoExamplesReadOnly

import "fmt"

/*
声明数组
Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：

var variable_name [SIZE] variable_type

var balance [10] float32

var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

float32 salary = balance[9]
*/

func mainfalse() {
	var n = [...]int{333, 2, 3, 4, 5}
	var i, j int
	for i = 0; i < 5; i++ {
		n[i] = n[i] + 100
	}
	for j = 0; j < 5; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
