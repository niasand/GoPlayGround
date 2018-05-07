/*
* @Author: Zhiwei Yang
* @Date:   2018-05-07 10:42:36
* @Last Modified by:   Zhiwei Yang
* @Last Modified time: 2018-05-07 10:52:04
 */
package main

import (
	"fmt"
)

func adder(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(adder(1, 3, 3))
	nums := []int{1, 2, 3, 4}
	fmt.Println(adder(nums...))
}
