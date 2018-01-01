/*
* @Author: jerry
* @Date:   2017-10-02 14:54:03
* @Last Modified by:   jerry
* @Last Modified time: 2017-10-02 14:58:36
 */

package main

import (
	"fmt"
)

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is %d\n", x)
	arr3 := []int{7, 9, 3, 5, 1}
	xx := min(arr3...)
	fmt.Printf("The minimum in the array arr is:%d\n", xx)

}

func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
