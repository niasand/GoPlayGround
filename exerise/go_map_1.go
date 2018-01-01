/*
* @Author: jerry
* @Date:   2017-09-22 11:44:41
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-22 11:48:26
 */
package main

import "fmt"

func main() {
	CountrycapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	fmt.Println("原始 map")

	for country := range CountrycapitalMap {
		fmt.Println("Capital of ", country, "is", CountrycapitalMap[country])
	}

	/* 删除元素 */
	delete(CountrycapitalMap, "France")
	fmt.Println("Entry for France is deleted")
	fmt.Println("删除元素后 map")

	for country := range CountrycapitalMap {
		fmt.Println("Capital", country, "is", CountrycapitalMap[country])
	}
}
