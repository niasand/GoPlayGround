/*
* @Author: jerry
* @Date:   2017-09-26 14:31:12
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-26 14:37:58
 */

package main

import "fmt"

func main() {
	var i int = 5
	for i >= 0 {
		i -= 1
		fmt.Printf("The variable i now is %d\n", i)
	}
}
