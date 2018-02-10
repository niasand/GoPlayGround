/*
* @Author: jerry
* @Date:   2018-01-22 14:43:36
* @Last Modified by:   jerry
* @Last Modified time: 2018-01-23 14:48:05
 */
package main

import "github.com/codeskyblue/go-sh"

func main() {
	sh.Command("echo", "hello").Run()
}
