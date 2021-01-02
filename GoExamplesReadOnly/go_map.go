/*
* @Author: jerry
* @Date:   2017-09-22 10:52:48
* @Last Modified by:   jerry
* @Last Modified time: 2017-09-22 11:43:38
 */

/* 声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type
*/

/* 使用 make 函数
map_variable := make(map[key_data_type]value_data_type)
如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
*/

package GoExamplesReadOnly

import "fmt"

func mainfalse() {
	var countryCapticalMap map[string]string
	countryCapticalMap = make(map[string]string)
	countryCapticalMap["France"] = "Paris"
	countryCapticalMap["Italy"] = "Rome"
	countryCapticalMap["Japan"] = "Tokyo"
	countryCapticalMap["India"] = "New Delhi"
	fmt.Println(countryCapticalMap)

	capital, ok := countryCapticalMap["Chinese"]
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of chinese is not present")
	}
}
