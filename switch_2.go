/*
* @Author: jerry
* @Date:   2017-09-21 15:49:01
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-21 15:53:46
 */
package main

import "fmt"

func main() {
	var grade string = "B"
	var marks int = 80

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀\n")
	case grade == "B":
		fmt.Printf("good\n")
	default:
		fmt.Printf("bad\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}
