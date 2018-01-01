/*
* @Author: jerry
* @Date:   2017-09-23 20:37:53
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-23 20:40:11
 */
package main

import "fmt"

func main() {
	i := 2
	switch {
	case i == 1:
		fmt.Println("i is 1")
	case i < 1:
		fmt.Println("i is less than 1 ")
	case i > 2:
		fmt.Println("i is more 1")
	default:
		fmt.Println("....")
	}
}
