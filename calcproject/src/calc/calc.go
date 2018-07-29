/*
 * @Script: calc.go
 * @Author: Zhiwei.Yang
 * @Email: tencrance@gmail.com
 * @Create At: 2018-07-29 10:06:12
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-29 15:12:31
 * @Description: This is description.
 */

package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USEAGE: cale command [arguments] ... ")
	fmt.Println("\nThe commands are \n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {

	args := os.Args[1:]
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	switch args[0] {
	case "add":
		if len(args) != 3 {
			fmt.Println("USAGE: calc add <integer1><intger2>")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		v2, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><intger2>")
			return
		}

		ret := simplemath.Add(v1, v2)
		fmt.Println("Result:", ret)

	case "sqrt":
		if len(args) != 2 {
			fmt.Println("USAGE: calc sqrt <intger> ")
			return
		}
		v, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <intger> ")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result:", ret)

	default:
		Usage()
	}
}
