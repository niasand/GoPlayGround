/*
* @Author: jerry
* @Date:   2017-09-23 20:43:05
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-23 20:46:13
 */
package main

import "fmt"

func main() {
	i := 3
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2, 3, 4:
		fmt.Println("i is 2")
	default:
		fmt.Println("------")
	}
}
