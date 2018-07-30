/*
 * @Script: noname.go
 * @Author: Zhiwei.Yang
 * @Email: tencrance@gmail.com
 * @Create At: 2018-07-30 21:56:49
 * @Last Modified By: Zhiwei.Yang
 * @Last Modified At: 2018-07-30 22:05:26
 * @Description: This is description.
 */

package main

import "fmt"

func main() {
	j := 5
	a := func() func() {
		i := 10
		return func() {
			fmt.Printf("i, j: %d ,%d\n", i, j)
		}
	}()
	a()

	j *= 2
	a()
}
