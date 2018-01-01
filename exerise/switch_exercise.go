/*
* @Author: jerry
* @Date:   2017-09-23 20:47:23
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-23 20:53:51
 */
package main

import "fmt"

func Season(month int) {
	switch month {
	case 1, 2, 3:
		fmt.Println("Spring..")
	case 4, 5, 6:
		fmt.Println("Summer..")
	case 7, 8, 9:
		fmt.Println("autumn")
	case 10, 11, 12:
		fmt.Println("Winter..")
	default:
		fmt.Println("Sorry! Can not find the season for this month")
	}
}

func main() {

	m := 10
	Season(m)
	switch m {
	case 1, 2, 3:
		fmt.Println("Spring..")
	case 4, 5, 6:
		fmt.Println("Summer..")
	case 7, 8, 9:
		fmt.Println("autumn")
	case 10, 11, 12:
		fmt.Println("Winter..")
	default:
		fmt.Println("Sorry! Can not find the season for this month")
	}
}
