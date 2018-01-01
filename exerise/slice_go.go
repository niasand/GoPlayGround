/*
* @Author: jerry
* @Date:   2017-09-21 20:33:21
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-22 10:02:03
 */
package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {
	// var numbers = make([]int, 3, 5) //make([]T, length, capacity)
	numbers := []int{0, 1, 2}
	numbers = append(numbers, 3, 4, 5, 6)
	numbers1 := make([]int, 0, 5)
	numbers2 := make([]int, len(numbers), (cap(numbers))*2)

	// var numbers []int

	if numbers == nil {
		fmt.Printf("切片是空的\n")
	} else {
		printSlice(numbers1)
	}
	fmt.Println(numbers)
	fmt.Println(numbers1)
	fmt.Println(numbers2)
}
