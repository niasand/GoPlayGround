/*
* @Author: jerry
* @Date:   2017-09-21 18:39:22
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-21 18:40:53
 */
package main

import "fmt"

func main() {
	var a = [5][2]int{{0, 0}, {1, 2}, {3, 4}, {5, 6}, {7, 8}}
	var i, j int
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
