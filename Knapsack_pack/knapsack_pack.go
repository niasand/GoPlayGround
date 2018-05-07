/*
* @Author: jerry
* @Date:   2017-09-28 18:21:16
* @Last Modified by:   Zhiwei Yang
* @Last Modified time: 2018-05-02 15:12:26
 */
package main

import (
	"fmt"
)

func getValue(W []int, V []int, MAX int, i int) (int, int) {
	var notiV, notW int
	if i > 1 {
		if W[i-1] > MAX {
			notiV, notW = getValue(W, V, MAX, i-1)
			return notiV, notW
		} else {
			puti_in_V, puti_in_W := getValue(W, V, MAX-W[i-1], i-1)
			if (puti_in_V + V[i-1]) > notiV {
				return (puti_in_V + V[i-1]), puti_in_W
			} else {
				return notiV, notW
			}
		}
	} else {
		if W[0] > MAX {
			return 0, MAX
		} else {
			return V[0], MAX - W[0]
		}
	}

}

func main() {
	V := []int{10, 20, 30, 40}
	W := []int{5, 4, 6, 2}
	MAX := 9
	fmt.Println(getValue(W, V, MAX, 4))

}
