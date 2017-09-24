/*
* @Author: jerry
* @Date:   2017-09-23 21:17:01
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-23 21:35:11
 */
package main

import "fmt"

func cal() {
	var i, j int
	for i = 0; i < 15; i++ {
		j += 1
		fmt.Println(j)
	}
}

func printG() {
	var i int
	var g string
	for i = 1; i < 25; i++ {
		g += "G"
		fmt.Printf("%s\n", g)
	}
}

func print_rectangle() {
	var i, j int
	// var m, k string
	for i = 0; i < 10; i++ {
		for j = 0; j < 20; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}

func main() {
	// cal()
	// printG()
	print_rectangle()
}
